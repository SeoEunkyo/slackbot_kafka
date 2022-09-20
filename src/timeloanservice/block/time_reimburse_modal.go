package block

import "github.com/slack-go/slack"

func GenerateTimeReimburseModal() slack.ModalViewRequest {
	// Create a ModalViewRequest with a header and two inputs
	titleText := slack.NewTextBlockObject("plain_text", "로동시간 저축 시전", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	headerText := slack.NewTextBlockObject("mrkdwn", "동무의 로동시간을 저축, 상환합니다.(최대 +-4시간, 하루 1시간이내)", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	divSection1 := slack.NewDividerBlock()
	startTimeLabel := slack.NewTextBlockObject("plain_text", "출근시간을 입력하세요", false, false)
	startTimeElement := slack.NewTimePickerBlockElement("startTime")
	startTimeElement.InitialTime = "10:00"
	startTimeBlock := slack.NewInputBlock("Start_Time", startTimeLabel, nil, startTimeElement)

	endTimeLabel := slack.NewTextBlockObject("plain_text", "퇴근시간을 입력하세요", false, false)
	endTimeElement := slack.NewTimePickerBlockElement("endTime")
	endTimeElement.InitialTime = "20:00"
	endTimeBlock := slack.NewInputBlock("End_Time", endTimeLabel, nil, endTimeElement)

	divSection2 := slack.NewDividerBlock()
	breakTimeLabel := slack.NewTextBlockObject("plain_text", "저녁식사시간(분)", false, false)
	breakTimeElement := slack.NewPlainTextInputBlockElement(breakTimeLabel, "breakTime")
	breakTimeElement.InitialValue = "0"
	breakTimeBlock := slack.NewInputBlock("BreakTime", breakTimeLabel, nil, breakTimeElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			divSection1,
			startTimeBlock,
			endTimeBlock,
			divSection2,
			breakTimeBlock,
		},
	}
	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	modalRequest.CallbackID = "time_reimburse"

	return modalRequest
}
