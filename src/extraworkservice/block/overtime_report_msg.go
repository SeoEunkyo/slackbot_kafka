package block

import (
	"fmt"

	"github.com/slack-go/slack"
)

func GenerateOvertimeReportMsg(applicant string, workTime int, breakTime int, thisWeekOvertime int) slack.MsgOption {

	headerText := slack.NewTextBlockObject("mrkdwn", "*<https://www.notion.so/teamdaydreamlab/424b73582e0c4ecf9f042ce525c88213|야근 보고.>*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	msg := fmt.Sprintf("%s 님이 금일 `%d시간 %d분 (식사시간:%d분)` 연장근무 하였습니다. \n", applicant, workTime/60, workTime%60, breakTime)
	msg += fmt.Sprintf("금주 야근시간은(총) `%d시간 %d분` 입니다.", thisWeekOvertime/60, thisWeekOvertime%60)
	bodyText := slack.NewTextBlockObject("mrkdwn", msg, false, false)
	bodySection := slack.NewSectionBlock(bodyText, nil, nil)

	return slack.MsgOptionBlocks(headerSection,
		bodySection)

}
