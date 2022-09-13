package rest

import (
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/extraworkservice/block"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/slack-go/slack"
	"net/http"
)

type SlashHandler struct {
	slackToken string
	dbhandler  persistence.DatabaseHandler
}

func NewSlashHandler(datahandler persistence.DatabaseHandler, token string) *SlashHandler {
	return &SlashHandler{
		dbhandler:  datahandler,
		slackToken: token,
	}
}

func (h *SlashHandler) ListenSlash(w http.ResponseWriter, r *http.Request) {
	err := verifySigningSecret(r)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	switch s.Command {
	case "/야근신청":
		api := slack.New(h.slackToken)
		modalRequest := block.GenerateOvertimeRequestModal()
		_, err = api.OpenView(s.TriggerID, modalRequest)
		if err != nil {
			fmt.Printf("Error opening view: %s", err)
		}
	case "/야근보고":
		api := slack.New(h.slackToken)
		approvedWorks := h.dbhandler.GetOvertimeCompleteYet()
		var modalRequest slack.ModalViewRequest
		if len(approvedWorks) > 0 {
			modalRequest = block.GenerateOvertimeReportModal(approvedWorks)
		} else {
			modalRequest = block.GenerateWarningModal("승인된 연장근무 내역이 없습니다. \n 신청을 먼저 진행하여 주세요!!")
		}

		_, err = api.OpenView(s.TriggerID, modalRequest)
		if err != nil {
			fmt.Printf("Error opening view: %s", err)
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
