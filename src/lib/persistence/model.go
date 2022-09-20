package persistence

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OvertimeRequest struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Applicant string
	Reason    string
	Date      string
	ApproveYN bool
}

type TimeBankRequest struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserName  string             `bson:"user_name" json:"user_name"`
	StartTime string             `bson:"start_time" json:"start_time"`
	EndTime   string             `bson:"end_time" json:"end_time"`
	BreakTime string             `bson:"break_time" json:"break_time"`
	Gap       int                `bson:"gap" json:"gap"`
	Amount    int                `bson:"amount" json:"amount"`
	Sum       int                `bson:"sum" json:"sum"`
	LeaveType string             `bson:"leave_type" json:"leave_type"` // 반차, 연차, 반반차
}
