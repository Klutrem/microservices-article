package kafka

import (
	"context"
	"fmt"
	"main/internal/config"
	"main/pkg"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type KafkaClient struct {
	env           config.Env
	client        sarama.Client
	producer      sarama.AsyncProducer
	consumerGroup sarama.ConsumerGroup
	logger        pkg.Logger
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
	conf.ClientID = uuid.NewString()
	conf.Consumer.Return.Errors = true

	client, err := sarama.NewClient([]string{addr}, conf)
	if err != nil {
		logger.Error(err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		logger.Error(err)
	}

	group, err := sarama.NewConsumerGroupFromClient(env.KafkaGroup, client)
	if err != nil {
		logger.Error(err)
	}
	go func() {
		for err := range group.Errors() {
			logger.Error(err)
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
	for {
		err := cl.consumerGroup.Consume(context.Background(), topics, handler)
		if err != nil {
			cl.logger.Error(err)
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
