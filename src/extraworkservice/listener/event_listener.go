package listener

import (
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/contracts"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/msgqueue"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence"
	"log"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
	SlackToken    string
}

func (p *EventProcessor) ProcessEvents() {
	log.Println("listening or events")

	received, errors, err := p.EventListener.Listen("slack_bot")

	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			fmt.Printf("got event %T: %s\n", evt, evt)
			p.handleEvent(evt)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.CallbackEvent:
		log.Printf("event %s created: %s", e.EventName(), e)

		switch e.CallbackId {
		case "overtime_request":
			p.OvertimeRequest(e.Payload)
		case "overtime_report":

		default:
			// button actionHandler

		}

	default:
		log.Printf("unknown event type: %T", e)
	}
}
