package block

import (
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/slack-go/slack"
)

type optionBlockValue struct {
	key   string
	value string
}

func createOptionBlockObjectsWithValue(options []optionBlockValue) []*slack.OptionBlockObject {
	optionBlockObjects := make([]*slack.OptionBlockObject, 0, len(options))

	for _, o := range options {
		optionText := slack.NewTextBlockObject(slack.PlainTextType, o.value, false, false)
		optionBlockObjects = append(optionBlockObjects, slack.NewOptionBlockObject(o.key, optionText, nil))
	}
	return optionBlockObjects
}

func GenerateOvertimeReportModal(approvedWorks []persistence.OvertimeRequest) slack.ModalViewRequest {

	titleText := slack.NewTextBlockObject("plain_text", "연장근무 보고", false, false)
	closeText := slack.NewTextBlockObject("plain_text", "Close", false, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	selectLabel := slack.NewTextBlockObject(slack.PlainTextType, "승인 리스트", false, false)

	selectContext := make([]optionBlockValue, 0)
	for _, request := range approvedWorks {
		minLength := func(length int) int {
			if length > 10 {
				return 10
			}
			return length
		}(len(request.Reason))

		context := request.Applicant + " - " + request.Date + "	[" + request.Reason[:minLength] + "....]"
		selectContext = append(selectContext, optionBlockValue{key: request.ID.Hex(), value: context})
	}
	selectList := createOptionBlockObjectsWithValue(selectContext)
	selectElement := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, nil, "overtime_report", selectList...)
	selectElement.Placeholder = slack.NewTextBlockObject(slack.PlainTextType, "선택해주세요", false, false)
	selectBlock := slack.NewInputBlock("select", selectLabel, nil, selectElement)

	divSection1 := slack.NewDividerBlock()
	startTimeLabel := slack.NewTextBlockObject("plain_text", "시작시간을 입력하세요", false, false)
	startTimeElement := slack.NewTimePickerBlockElement("startTime")
	startTimeElement.InitialTime = "20:30"
	startTimeBlock := slack.NewInputBlock("Start_Time", startTimeLabel, nil, startTimeElement)

	endTimeLabel := slack.NewTextBlockObject("plain_text", "종료시간을 입력하세요", false, false)
	endTimeElement := slack.NewTimePickerBlockElement("endTime")
	endTimeElement.InitialTime = "22:30"
	endTimeBlock := slack.NewInputBlock("End_Time", endTimeLabel, nil, endTimeElement)

	divSection2 := slack.NewDividerBlock()
	breakTimeLabel := slack.NewTextBlockObject("plain_text", "식사시간(분)", false, false)
	breakTimeElement := slack.NewPlainTextInputBlockElement(breakTimeLabel, "breakTime")
	breakTimeElement.InitialValue = "30"
	breakTimeBlock := slack.NewInputBlock("BreakTime", breakTimeLabel, nil, breakTimeElement)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			selectBlock,
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
	modalRequest.CallbackID = "overtime_report"

	return modalRequest
}
