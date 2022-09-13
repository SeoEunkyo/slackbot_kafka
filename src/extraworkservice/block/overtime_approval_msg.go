package block

import "github.com/slack-go/slack"

func GenerateOvertimeApprovalMsg(displayName string, originBlocks slack.Blocks) slack.MsgOption {
	approvalText := slack.NewTextBlockObject("mrkdwn", displayName+" 님이 "+"`승인` 하였습니다. ", false, false)
	approvalSection := slack.NewSectionBlock(approvalText, nil, nil)

	// 승인시에 DB에 저장합니다.

	return slack.MsgOptionBlocks(originBlocks.BlockSet[0], originBlocks.BlockSet[1], approvalSection)
}
