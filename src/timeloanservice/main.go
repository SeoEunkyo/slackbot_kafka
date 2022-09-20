package main

import (
	"flag"
	"fmt"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/configuration"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/msgqueue/kafka"
	"github.com/SeoEunkyo/slackbot_kafka/src/lib/persistence/dblayer"
	"github.com/SeoEunkyo/slackbot_kafka/src/timeloanservice/listener"
	"github.com/SeoEunkyo/slackbot_kafka/src/timeloanservice/rest"
	"github.com/Shopify/sarama"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	confPath := flag.String("conf", "./config.json", "flag to set the path to the configuration json file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("config : ", config)

	//create kafka emitter
	conf := sarama.NewConfig()
	conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
	panicIfErr(err)
	//consumer 생성
	eventListener, err := kafka.NewKafkaEventListener(conn, []int32{})
	panicIfErr(err)

	dbhandler, _ := dblayer.NewPersistenceLayer(dblayer.DBTYPE(config.Databasetype), config.DBConnection)
	processor := listener.EventProcessor{eventListener, dbhandler, config.SlackToken}

	go processor.ProcessEvents()
	fmt.Println("restfullEndPoint :", config.RestfulEndpoint)
	rest.ServeAPI(config.RestfulEndpoint, config.SlackToken, dbhandler)

}
