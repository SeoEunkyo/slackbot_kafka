package listener

import (
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/extraworkservice/block"
	"github.com/slack-go/slack"
	"time"
)

// OvertimeReqApproval  overtime_req_approval
func (p *EventProcessor) OvertimeReqApproval(i slack.InteractionCallback) {
	api := slack.New(p.SlackToken)
	t := time.Now()
	userParam := slack.GetUserProfileParameters{}
	userParam.UserID = i.User.ID
	userInfo, _ := api.GetUserProfile(&userParam)

	msgBlock := block.GenerateOvertimeApprovalMsg(userInfo.DisplayName, i.Message.Msg.Blocks)

	// DB에 overtime정보 저장

	oid := i.ActionCallback.BlockActions[0].Value
	p.Database.ApproveOvertime(oid)

	_, _, _, err := api.UpdateMessage(i.User.ID, t.Format(time.RFC3339),
		msgBlock,
		slack.MsgOptionReplaceOriginal(i.ResponseURL))
	// h.dbhandler.GetOvertimeCompleteYet()
	if err != nil {
		fmt.Printf(err.Error())
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
