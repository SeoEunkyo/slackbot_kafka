package persistence

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OvertimeRequest struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Applicant    string
	Reason     string
	Date string
	ApproveYN bool
}