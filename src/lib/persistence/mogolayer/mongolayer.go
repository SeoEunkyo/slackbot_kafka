package mongolayer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB, Tables
const (
	DB       = "BmeksSlack"
	Overtime = "Overtime"
)

type MongoDBLayer struct {
	client *mongo.Client
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {

	// "mongodb://localhost:27017"
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))
	if err != nil {
		panic(err)
	}
	return &MongoDBLayer{
		client: c,
	}, err
}

func (mgolayer *MongoDBLayer) RequestOvertime(req persistence.OvertimeRequest) primitive.ObjectID {

	coll := mgolayer.client.Database(DB).Collection(Overtime)
	doc := bson.D{{"applicant", req.Applicant}, {"date", req.Date},
		{"doneYN", false}, {"reason", req.Reason},
		{"approveYN", req.ApproveYN}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	var returnOid primitive.ObjectID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		returnOid = oid
		return returnOid
	}
	return returnOid
}

func (mgolayer *MongoDBLayer) ApproveOvertime(oid string) {

	coll := mgolayer.client.Database(DB).Collection(Overtime)
	id, _ := primitive.ObjectIDFromHex(oid[10:34])
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"approveYN", true}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func (mgolayer *MongoDBLayer) ReportOvertimeWork(oidHex string, mins int) {

	coll := mgolayer.client.Database(DB).Collection(Overtime)
	id, _ := primitive.ObjectIDFromHex(oidHex)
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"doneYN", true}, {"mins", mins}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
func (mgolayer *MongoDBLayer) GetSumOfOverWorkInThisWeek(oidHex string) (int, string) {
	coll := mgolayer.client.Database(DB).Collection(Overtime)
	id, _ := primitive.ObjectIDFromHex(oidHex)
	filter := bson.D{{"_id", id}}
	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	timeC, _ := time.Parse("2006-01-02", fmt.Sprintf("%v", result["date"]))
	yearC, weekC := timeC.ISOWeek()
	fmt.Print(yearC, weekC)
	filter = bson.D{{"applicant", fmt.Sprintf("%v", result["applicant"])}, {"mins", bson.M{"$ne": nil}}}
	projection := bson.D{{"_id", 1}, {"mins", 2}, {"date", 3}, {"reason", 4}}
	opts := options.Find().SetProjection(projection)
	cursor, err := coll.Find(context.TODO(), filter, opts)
	var resultsD []bson.D

	if err = cursor.All(context.TODO(), &resultsD); err != nil {
		panic(err)
	}
	SumOfWork := 0
	for _, r := range resultsD {

		dateOfWork := fmt.Sprintf("%v", r[1].Value)
		timeT, _ := time.Parse("2006-01-02", dateOfWork)
		yearT, weekT := timeT.ISOWeek()
		if r[3].Value == nil {
			continue
		}
		mins := fmt.Sprintf("%v", r[3].Value)
		if yearT == yearC && weekC == weekT {
			min, _ := strconv.Atoi(mins)
			SumOfWork += min
		}
	}

	return SumOfWork, fmt.Sprintf("%v", result["applicant"])

}

func (mgolayer *MongoDBLayer) GetOvertimeCompleteYet() []persistence.OvertimeRequest {

	coll := mgolayer.client.Database(DB).Collection(Overtime)
	filter := bson.D{{"doneYN", false}, {"approveYN", true}}
	projection := bson.D{{"_id", 1}, {"applicant", 2}, {"date", 3}, {"reason", 4}}
	opts := options.Find().SetProjection(projection)
	cursor, err := coll.Find(context.TODO(), filter, opts)
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	returnValues := make([]persistence.OvertimeRequest, 0)
	for _, result := range results {
		fmt.Println(result[0].Value, result[1].Value, result[2].Value)
		fmt.Print(result[0].Value)
		hex := fmt.Sprintf("%v", result[0].Value)
		objectID, _ := primitive.ObjectIDFromHex(hex[10:34])
		newItem := persistence.OvertimeRequest{ID: objectID,
			Applicant: fmt.Sprintf("%v", result[1].Value),
			Date:      fmt.Sprintf("%v", result[2].Value),
			Reason:    fmt.Sprintf("%v", result[3].Value)}
		returnValues = append(returnValues, newItem)
	}

	return returnValues

}
