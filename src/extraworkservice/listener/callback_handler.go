package listener

import (
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/extraworkservice/block"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/slack-go/slack"
	"strconv"
)

func (p *EventProcessor) OvertimeRequest(i slack.InteractionCallback) {
	api := slack.New(p.SlackToken)
	userParam := slack.GetUserProfileParameters{}
	userParam.UserID = i.User.ID
	userInfo, _ := api.GetUserProfile(&userParam)

	pickDate := i.View.State.Values["Pick Date"]["datePicker"].SelectedDate
	reason := i.View.State.Values["Reason"]["Reason"].Value

	overReq := persistence.OvertimeRequest{Applicant: userInfo.DisplayName,
		ApproveYN: false,
		Reason:    reason,
		Date:      pickDate,
	}
	oid := p.Database.RequestOvertime(overReq)
	msgBlock := block.GenerateOvertimeResponseMsg(userInfo.DisplayName, pickDate, reason, oid.String())

	_, _, err := api.PostMessage(i.User.ID, msgBlock)
	// _, _, err = api.PostMessage(i.User.ID, msgBlock ,options)
	if err != nil {
		fmt.Printf(err.Error())
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// OvertimeReport overtime_report
func (p *EventProcessor) OvertimeReport(i slack.InteractionCallback) {
	selectedOption := i.View.State.Values["select"]["overtime_report"].SelectedOption
	oidHex := selectedOption.Value

	startTimeSting := i.View.State.Values["Start_Time"]["startTime"].SelectedTime
	startTime := ConvertStringToIntTime(startTimeSting)

	endTimeString := i.View.State.Values["End_Time"]["endTime"].SelectedTime
	endTime := ConvertStringToIntTime(endTimeString)
	// gap := endTime - startTime

	if startTime.Hours > endTime.Hours {
		endTime.Hours += 24
	}
	breakTimeString := i.View.State.Values["BreakTime"]["breakTime"].Value
	breakTime, _ := strconv.Atoi(breakTimeString)
	gap := ((endTime.Hours*60 + endTime.Mins) - (startTime.Hours*60 + startTime.Mins)) - breakTime

	// 현재까지의 근무 데이터 필요 (현재까지 근무시간 표시  52시간 제도)
	thisWeekOvertimeInMins, applicant := p.Database.GetSumOfOverWorkInThisWeek(oidHex)
	// Db에 넣고
	p.Database.ReportOvertimeWork(oidHex, gap)
	// kafka 전송
	// 연차 관리 시스템의 사용을 예상한ㄷㅏ.

	// 메세지 출력
	// 금주의 야근 시간은 현재 몇시간 몇분입니다
	msg := block.GenerateOvertimeReportMsg(applicant, gap, breakTime, thisWeekOvertimeInMins+gap)
	api := slack.New(p.SlackToken)
	_, _, err := api.PostMessage(i.User.ID, msg)
	// _, _, err = api.PostMessage(i.User.ID, msgBlock ,options)
	if err != nil {
		fmt.Printf(err.Error())
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
