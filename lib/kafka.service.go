package lib

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/opentracing/opentracing-go/log"
)

type KafkaClient struct {
	env      Env
	consumer sarama.Consumer
	producer sarama.AsyncProducer
}

func NewKafkaClient(env Env) KafkaClient {
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
		env:      env,
		consumer: consumer,
		producer: producer,
	}
}

func (cl *KafkaClient) Consume(topic string) <-chan *sarama.ConsumerMessage {
	consumer, err := cl.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Error(err)
	}
	return consumer.Messages()
}

func (cl *KafkaClient) Send(topic string, message string) {
	value := sarama.StringEncoder(message)
	cl.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: value,
	}
}
