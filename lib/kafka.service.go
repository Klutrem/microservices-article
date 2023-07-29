package lib

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/opentracing/opentracing-go/log"
)

type KafkaClient struct {
	env           Env
	client        sarama.Client
	producer      sarama.AsyncProducer
	ConsumerGroup sarama.ConsumerGroup
}

type KafkaHandler interface {
	Handle(topic string, message []byte)
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
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

	group, err := sarama.NewConsumerGroupFromClient(env.KafkaGroup, client)
	if err != nil {
		log.Error(err)
	}
	return KafkaClient{
		env:           env,
		producer:      producer,
		ConsumerGroup: group,
		client:        client,
	}
}

func (cl *KafkaClient) Consume(handler KafkaHandler, topics []string) {
	consumer, err := sarama.NewConsumerGroupFromClient("test", cl.client)
	if err != nil {
		log.Error(err)
	}
	defer consumer.Close()

	for {
		consumer.Consume(context.Background(), topics[:], handler)
	}
}

func (cl *KafkaClient) Reply(topic string, message string) {
	value := sarama.StringEncoder(message)
	cl.producer.Input() <- &sarama.ProducerMessage{
		Topic: fmt.Sprint(topic, ".reply"),
		Value: value,
	}
}

func (cl *KafkaClient) Send(topic string, message string) {
	value := sarama.StringEncoder(message)
	cl.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: value,
	}
}
