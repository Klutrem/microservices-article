package lib

import (
	"fmt"
	"main/lib"

	"github.com/IBM/sarama"
	"github.com/opentracing/opentracing-go/log"
)

type KafkaClient struct {
	Env      lib.Env
	Consumer sarama.Consumer
	Producer sarama.AsyncProducer
}

func NewKafkaClient(env lib.Env) KafkaClient {
	addr := fmt.Sprint(env.BrokerHost, ":", env.BrokerPort)
	conf := sarama.NewConfig()
	client, err := sarama.NewClient([]string{addr}, conf)
	if err != nil {
		log.Error(err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Error(err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Error(err)
	}
	return KafkaClient{
		Env:      env,
		Consumer: consumer,
		Producer: producer,
	}
}
