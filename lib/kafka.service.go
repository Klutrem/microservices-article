package lib

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/opentracing/opentracing-go/log"
)

type KafkaClient struct {
	env           Env
	client        sarama.Client
	producer      sarama.SyncProducer
	reader        sarama.Consumer
	consumerGroup sarama.ConsumerGroup
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

	producer, err := sarama.NewSyncProducerFromClient(client)
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
		consumerGroup: group,
		client:        client,
	}
}

func (cl *KafkaClient) Consume(handler KafkaHandler, topics []string) {
	defer cl.consumerGroup.Close()

	for {
		cl.consumerGroup.Consume(context.Background(), topics[:], handler)
	}
}

func (cl *KafkaClient) Send(topic string, message []byte) (response []byte, err error) {
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	partition, offset, err := cl.producer.SendMessage(&msg)
	if err != nil {
		return nil, err
	}
	replyTopic := fmt.Sprintf(topic, ".reply")
	consumer, err := sarama.NewConsumerFromClient(cl.client)
	if err != nil {
		return nil, err
	}
	partitionConsumer, err := consumer.ConsumePartition(replyTopic, partition, offset)
	if err != nil {
		consumer.Close()
		return nil, err
	}
	select {
	case msg := <-partitionConsumer.Messages():
		response = msg.Value
	case <-time.After(time.Second * 60):
		consumer.Close()
		return nil, errors.New("timeout waiting for reply message")
	}
	consumer.Close()

	return response, nil

}
