package listener

import (
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/extraworkservice/block"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/slack-go/slack"
)

func (p *EventProcessor) OvertimeRequest(i slack.InteractionCallback) {
	api := slack.New(p.SlackToken)
	userParam := slack.GetUserProfileParameters{}
	userParam.UserID = i.User.ID
	userInfo, _ := api.GetUserProfile(&userParam)

	pickDate := i.View.State.Values["Pick Date"]["datePicker"].SelectedDate
	reason := i.View.State.Values["Reason"]["Reason"].Value

	overReq := persistence.OvertimeRequest{Applicant: userInfo.DisplayName,
		ApproveYN: false,
		Reason:    reason,
		Date:      pickDate,
	}
	oid := p.Database.RequestOvertime(overReq)
	msgBlock := block.GenerateOvertimeResponseMsg(userInfo.DisplayName, pickDate, reason, oid.String())

	_, _, err := api.PostMessage(i.User.ID, msgBlock)
	// _, _, err = api.PostMessage(i.User.ID, msgBlock ,options)
	if err != nil {
		fmt.Printf(err.Error())
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
