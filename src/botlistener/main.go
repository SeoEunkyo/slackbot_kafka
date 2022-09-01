package main

import (
	"flag"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/seoEunkyo/slackbot_kafka/src/botlistener/rest"
	"github.com/seoEunkyo/slackbot_kafka/src/lib/configuration"
	"github.com/seoEunkyo/slackbot_kafka/src/lib/msgqueue/kafka"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// slack에서 interactivity payload를 받아서,
	// kafka에 메시지를 발행

	// kafka connection
	// rest API for receive msg from slack
	confPath := flag.String("conf", "./config.json", "flag to set the path to the configuration json file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("config : ", config)

	//create kafka emitter
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
	panicIfErr(err)

	eventEmitter, err := kafka.NewKafkaEventEmitter(conn)
	panicIfErr(err)

	rest.ServeAPI(config.RestfulEndpoint, eventEmitter)

}
