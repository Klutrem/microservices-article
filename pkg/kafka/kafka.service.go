package kafka

import (
	"context"
	"errors"
	"fmt"
	"main/internal/config"
	"main/pkg"
	"time"

	"github.com/IBM/sarama"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
)

type KafkaClient struct {
	env           config.Env
	client        sarama.Client
	producer      sarama.AsyncProducer
	consumerGroup sarama.ConsumerGroup
}

type KafkaHandler interface {
	Handle(topic string, message sarama.ConsumerMessage)
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
}

func NewKafkaClient(env config.Env, logger pkg.Logger) KafkaClient {
	sarama.Logger = zap.NewStdLog(logger.Desugar().Named("Kafka"))
	addr := fmt.Sprint(env.BrokerHost, ":", env.BrokerPort)
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
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
	go func() {
		for err := range group.Errors() {
			log.Error(err)
		}
	}()

	return KafkaClient{
		env:           env,
		producer:      producer,
		consumerGroup: group,
		client:        client,
	}
}

func (cl KafkaClient) Consume(handler KafkaHandler, topics []string) {
	defer cl.consumerGroup.Close()
	for {
		err := cl.consumerGroup.Consume(context.Background(), topics, handler)
		if err != nil {
			log.Error(err)
		}
	}
}

func (cl *KafkaClient) Send(topic string, message []byte, correlationID string) {
	value := sarama.StringEncoder(message)
	cl.producer.Input() <- &sarama.ProducerMessage{
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("kafka_nest-is-disposed"),
				Value: []byte("true"),
			},

			{
				Key:   []byte("kafka_correlationId"),
				Value: []byte(correlationID),
			},
		},
		Topic: topic,
		Value: value,
	}
}

func (cl *KafkaClient) SendWithReply(topic string, message []byte) (response []byte, err error) {
	value := sarama.StringEncoder(message)
	
	replyTopic := topic + ".reply"
	cl.producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: value,
		Headers: []sarama.RecordHeader{
			{Key: []byte("kafka_replyTopic"), Value: []byte(replyTopic)},
		},
	}
	
	consumer, err := sarama.NewConsumerFromClient(cl.client)
	if err != nil {
		return nil, err
	}
	defer consumer.Close()
	offset := sarama.OffsetNewest
	partitionConsumer, err := consumer.ConsumePartition(replyTopic, 0, offset)
	if err != nil {
		consumer.Close()
		return nil, err
	}
	defer partitionConsumer.Close()

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
