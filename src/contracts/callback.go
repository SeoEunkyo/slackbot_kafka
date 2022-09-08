package contracts

import "github.com/slack-go/slack"

type CallbackEvent struct {
	CallbackId string                    `json:"callbackId"`
	Payload    slack.InteractionCallback `json:"payload"`
}

func (c *CallbackEvent) EventName() string {
	return "callback"
}
