package block

import "github.com/slack-go/slack"

func GenerateOvertimeDenyMsg(displayName string) slack.SectionBlock {
	approvalText := slack.NewTextBlockObject("mrkdwn", displayName+" 님이 "+"`거절` 하였습니다. ", false, false)
	approvalSection := slack.NewSectionBlock(approvalText, nil, nil)
	return *approvalSection
}
