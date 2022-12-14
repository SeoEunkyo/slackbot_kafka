package listener

import (
	"github.com/slack-go/slack"
	"strconv"
	"strings"
)

type Times struct {
	Hours int
	Mins  int
}

func ConvertStringToIntTime(time string) *Times {
	times := strings.Split(time, ":")
	iHour, _ := strconv.Atoi(times[0])
	iMins, _ := strconv.Atoi(times[1])
	return &Times{Hours: iHour,
		Mins: iMins}

}

func EventHandler(p *EventProcessor, i slack.InteractionCallback) {

	switch i.View.CallbackID {
	case "overtime_request":
		p.OvertimeRequest(i)
	case "overtime_report":
		p.OvertimeReport(i)
	}
	if len(i.ActionCallback.BlockActions) < 1 {
		return
	}

	//actionEvent 처리
	switch i.ActionCallback.BlockActions[0].ActionID {
	case "overtime_req_approval":
		p.OvertimeReqApproval(i)
	case "":
	}
}
