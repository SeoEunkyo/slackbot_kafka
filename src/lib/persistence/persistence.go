package persistence

import "go.mongodb.org/mongo-driver/bson/primitive"

type DatabaseHandler interface {
	RequestOvertime(OvertimeRequest) primitive.ObjectID
	ApproveOvertime(oid string)
	GetOvertimeCompleteYet() []OvertimeRequest
	ReportOvertimeWork(oidHex string, mins int)
	GetSumOfOverWorkInThisWeek(oidHex string) (int, string)
	GetSavedTotalMin(userName string) int
	SavedTime(request *TimeBankRequest) (primitive.ObjectID, error)
}
