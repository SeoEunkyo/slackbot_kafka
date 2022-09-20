package block

import (
	model "github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/slack-go/slack"
	"strconv"
)

func GenerateTimeReimburseMsg(request *model.TimeBankRequest) slack.MsgOption {
	headerText := slack.NewTextBlockObject("mrkdwn", "\n*<https://www.notion.so/teamdaydreamlab/424b73582e0c4ecf9f042ce525c88213|"+request.UserName+" - 시간을 소비합니다.>*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	typeField := slack.NewTextBlockObject("mrkdwn", "*Type:*\n노동시간 상환", false, false)
	whenField := slack.NewTextBlockObject("mrkdwn", "*소비시간(분):*\n"+strconv.Itoa(request.Amount), false, false)
	// lastUpdateField := slack.NewTextBlockObject("mrkdwn", "*Request date:*\nMar 10, 2015 (3 years, 5 months)", false, false)
	reasonField := slack.NewTextBlockObject("mrkdwn", "*합계(분):*\n"+strconv.Itoa(request.Amount+request.Sum), false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)
	fieldSlice = append(fieldSlice, whenField)
	// fieldSlice = append(fieldSlice, lastUpdateField)
	fieldSlice = append(fieldSlice, reasonField)
	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)
	// Approve and Deny Buttons

	return slack.MsgOptionBlocks(
		headerSection,
		fieldsSection)
}
