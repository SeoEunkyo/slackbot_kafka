package dblayer

import (
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	mongolayer "github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence/mogolayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
