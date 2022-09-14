package rest

import (
	"encoding/json"
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/contracts"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/msgqueue"
	"github.com/slack-go/slack"
	"net/http"
)

type interactionHandler struct {
	eventEmitter msgqueue.EventEmitter
}

func NewInteractionsHandler(eventEmitter msgqueue.EventEmitter) *interactionHandler {
	return &interactionHandler{eventEmitter}
}

func (h *interactionHandler) ListenMsg(w http.ResponseWriter, r *http.Request) {
	err := verifySigningSecret(r)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var i slack.InteractionCallback
	err = json.Unmarshal([]byte(r.FormValue("payload")), &i)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//msg를 확인하고 kafka에 메세지를 publish
	msg := contracts.CallbackEvent{
		CallbackId: i.View.CallbackID,
		Payload:    i,
	}

	h.eventEmitter.Emit(&msg)

	return
}
