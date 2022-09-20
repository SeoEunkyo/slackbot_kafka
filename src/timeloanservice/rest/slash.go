package rest

import (
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"github.com/SeoEunkyo/slackbot_kafka/src/timeloanservice/block"
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
	case "/시간대출":
		api := slack.New(h.slackToken)
		modalRequest := block.GenerateTimeLoanModal()
		_, err = api.OpenView(s.TriggerID, modalRequest)
		if err != nil {
			fmt.Printf("Error opening view: %s", err)
		}
	case "/시간상환":
		api := slack.New(h.slackToken)
		modalRequest := block.GenerateTimeReimburseModal()
		_, err = api.OpenView(s.TriggerID, modalRequest)
		if err != nil {
			fmt.Printf("Error opening view: %s", err)
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
