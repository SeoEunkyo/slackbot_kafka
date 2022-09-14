package listener

import "github.com/slack-go/slack"

func EventHandler(p *EventProcessor, i slack.InteractionCallback) {

	switch i.View.CallbackID {
	case "overtime_request":
		p.OvertimeRequest(i)
	case "overtime_report":

	}
	if len(i.ActionCallback.BlockActions) < 1 {
		return
	}

	//actionEvent 처리
	switch i.ActionCallback.BlockActions[0].ActionID {
	case "overtime_req_approval":
		p.OvertimeReqApproval(i)

	}
}
