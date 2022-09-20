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
	case "time_loan":
		p.TimeLoanReq(i)
	case "time_reimburse":
		p.TimeReimburseReq(i)
	}

}
