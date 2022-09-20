package listener

import (
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/SeoEunkyo/slackbot_kafka/src/timeloanservice/block"
	"github.com/slack-go/slack"
	"strconv"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *EventProcessor) TimeLoanReq(i slack.InteractionCallback) {
	api := slack.New(p.SlackToken)
	userParam := slack.GetUserProfileParameters{}
	userParam.UserID = i.User.ID
	userInfo, _ := api.GetUserProfile(&userParam)

	startTimeSting := i.View.State.Values["Start_Time"]["startTime"].SelectedTime
	startTime := ConvertStringToIntTime(startTimeSting)

	endTimeString := i.View.State.Values["End_Time"]["endTime"].SelectedTime
	endTime := ConvertStringToIntTime(endTimeString)

	if startTime.Hours > endTime.Hours {
		endTime.Hours += 24
	}
	// 연차 반차 반반차 계산식들어가야함.
	breakTimeString := i.View.State.Values["BreakTime"]["breakTime"].Value
	breakTime, _ := strconv.Atoi(breakTimeString)
	gap := ((endTime.Hours*60 + endTime.Mins) - (startTime.Hours*60 + startTime.Mins)) - breakTime
	amount := gap - (9 * 60) // 하루에 연차와 시간저장을 같이 사용불가하다
	// 이전의 총시간을 가지고온다

	//p.Database.GetSavedTime()

	req := &persistence.TimeBankRequest{UserName: userInfo.DisplayName,
		StartTime: startTimeSting,
		EndTime:   endTimeString,
		BreakTime: breakTimeString,
		Gap:       gap,
		Amount:    amount,
	}
	req.Sum = p.Database.GetSavedTotalMin(userInfo.DisplayName)
	_, err := p.Database.SavedTime(req)
	panicIfErr(err)
	msgBlock := block.GenerateTimeLoanMsg(req)

	_, _, err = api.PostMessage(i.User.ID, msgBlock)
	// _, _, err = api.PostMessage(i.User.ID, msgBlock ,options)
	panicIfErr(err)
}

// TimeReimburseReq overtime_report
func (p *EventProcessor) TimeReimburseReq(i slack.InteractionCallback) {

	api := slack.New(p.SlackToken)
	userParam := slack.GetUserProfileParameters{}
	userParam.UserID = i.User.ID
	userInfo, _ := api.GetUserProfile(&userParam)

	startTimeSting := i.View.State.Values["Start_Time"]["startTime"].SelectedTime
	startTime := ConvertStringToIntTime(startTimeSting)

	endTimeString := i.View.State.Values["End_Time"]["endTime"].SelectedTime
	endTime := ConvertStringToIntTime(endTimeString)

	if startTime.Hours > endTime.Hours {
		endTime.Hours += 24
	}
	// 연차 반차 반반차 계산식들어가야함.
	breakTimeString := i.View.State.Values["BreakTime"]["breakTime"].Value
	breakTime, _ := strconv.Atoi(breakTimeString)
	gap := ((endTime.Hours*60 + endTime.Mins) - (startTime.Hours*60 + startTime.Mins)) - breakTime
	amount := gap - (9 * 60) // 하루에 연차와 시간저장을 같이 사용불가하다
	// 이전의 총시간을 가지고온다

	//p.Database.GetSavedTime()

	req := &persistence.TimeBankRequest{UserName: userInfo.DisplayName,
		StartTime: startTimeSting,
		EndTime:   endTimeString,
		BreakTime: breakTimeString,
		Gap:       gap,
		Amount:    amount,
	}
	req.Sum = p.Database.GetSavedTotalMin(userInfo.DisplayName)
	_, err := p.Database.SavedTime(req)
	panicIfErr(err)
	msgBlock := block.GenerateTimeReimburseMsg(req)

	_, _, err = api.PostMessage(i.User.ID, msgBlock)
	// _, _, err = api.PostMessage(i.User.ID, msgBlock ,options)
	panicIfErr(err)

}
