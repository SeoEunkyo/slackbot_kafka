package block

import "github.com/slack-go/slack"

func GenerateOvertimeResponseMsg(applicant string, pickDate string, reason string, oid string) slack.MsgOption {

	headerText := slack.NewTextBlockObject("mrkdwn", "new request :\n*<https://www.notion.so/teamdaydreamlab/424b73582e0c4ecf9f042ce525c88213|"+applicant+" - 야근을 신청합니다.>*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	typeField := slack.NewTextBlockObject("mrkdwn", "*Type:*\n야근 신청", false, false)
	whenField := slack.NewTextBlockObject("mrkdwn", "*When:*\n"+pickDate, false, false)
	// lastUpdateField := slack.NewTextBlockObject("mrkdwn", "*Request date:*\nMar 10, 2015 (3 years, 5 months)", false, false)
	reasonField := slack.NewTextBlockObject("mrkdwn", "*Reason:*\n"+reason, false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)
	fieldSlice = append(fieldSlice, whenField)
	// fieldSlice = append(fieldSlice, lastUpdateField)
	fieldSlice = append(fieldSlice, reasonField)
	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)
	// Approve and Deny Buttons
	approveBtnTxt := slack.NewTextBlockObject("plain_text", "Approve", false, false)
	approveBtn := slack.NewButtonBlockElement("overtime_req_approval", oid, approveBtnTxt)
	approveBtn.Style = slack.StylePrimary
	denyBtnTxt := slack.NewTextBlockObject("plain_text", "Deny", false, false)
	denyBtn := slack.NewButtonBlockElement("overtime_req_deny", oid, denyBtnTxt)
	denyBtn.Style = slack.StyleDanger
	actionBlock := slack.NewActionBlock("", approveBtn, denyBtn)
	return slack.MsgOptionBlocks(headerSection,
		fieldsSection,
		actionBlock)

}
