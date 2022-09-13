package block

import "github.com/slack-go/slack"

func GenerateWarningModal(msg string) slack.ModalViewRequest {

	titleText := slack.NewTextBlockObject("plain_text", "경고", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", msg, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
		},
	}
	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Blocks = blocks

	return modalRequest
}
