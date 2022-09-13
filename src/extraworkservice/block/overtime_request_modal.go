package block

import "github.com/slack-go/slack"

func GenerateOvertimeRequestModal() slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject("plain_text", "연장 근무 신청", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", "연장 근무 사유를 입력해 주세요.", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	dateLabel := slack.NewTextBlockObject("plain_text", "날짜", false, false)
	dateElement := slack.NewDatePickerBlockElement("datePicker")
	datePickerBlock := slack.NewInputBlock("Pick Date", dateLabel, nil, dateElement)

	reasonText := slack.NewTextBlockObject("plain_text", "사유", false, false)
	// lastNameHint := slack.NewTextBlockObject("plain_text", "Last Name Hint", false, false)
	reasonPlaceholder := slack.NewTextBlockObject("plain_text", "Enter your reason here", false, false)
	reasonElement := slack.NewPlainTextInputBlockElement(reasonPlaceholder, "Reason")
	reasonElement.Multiline = true
	reasonBlock := slack.NewInputBlock("Reason", reasonText, nil, reasonElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			datePickerBlock,
			reasonBlock,
		},
	}
	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	modalRequest.CallbackID = "overtime_request"

	return modalRequest
}
