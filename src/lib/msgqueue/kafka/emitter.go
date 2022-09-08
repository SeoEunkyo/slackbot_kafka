package kafka

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SeoEunkyo/slackbot_kafka/src/lib/helper/kafka"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/msgqueue"
	"github.com/Shopify/sarama"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

type messageEnvelope struct {
	EventName string      `json:"eventName"`
	Context   interface{} `json:"context"`
}

func NewKafkaEventEmitterFromEnvironment(connectionStrings []string) (msgqueue.EventEmitter, error) {
	brokers := connectionStrings

	if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
		brokers = strings.Split(brokerList, ",")
	}

	client := <-kafka.RetryConnect(brokers, 5*time.Second)
	return NewKafkaEventEmitter(client)
}
func NewKafkaEventEmitter(client sarama.Client) (msgqueue.EventEmitter, error) {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	emitter := kafkaEventEmitter{
		producer: producer,
	}

	return &emitter, nil
}

func (k *kafkaEventEmitter) Emit(evt msgqueue.Event) error {
	jsonBody, err := json.Marshal(messageEnvelope{
		evt.EventName(),
		evt,
	})
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "slack_bot",
		Value: sarama.ByteEncoder(jsonBody),
	}

	log.Printf("published message with topic %s: %v", evt.EventName(), jsonBody)
	_, _, err = k.producer.SendMessage(msg)

	return err
}
