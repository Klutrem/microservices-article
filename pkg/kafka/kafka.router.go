package kafka

import (
	"errors"
	"main/pkg"

	"github.com/IBM/sarama"
)

type KafkaReplyHandlerFunc func(message KafkaMessage) (response []byte, err error)

type KafkaRouter struct {
	replyHandlers map[string]KafkaReplyHandlerFunc
	kafkaClient   KafkaClient
	logger        pkg.Logger
}

func NewKafkaRouter(logger pkg.Logger, kafkaClient KafkaClient) KafkaRouter {
	return KafkaRouter{
		replyHandlers: make(map[string]KafkaReplyHandlerFunc),
		logger:        logger,
		kafkaClient:   kafkaClient,
	}
}

func findHeader(headers []*sarama.RecordHeader, key string) (string, error) {
	for _, header := range headers {
		if string(header.Key) == key {
			return string(header.Value), nil
		}
	}
	return "", errors.New("header not found")
}

func (kr *KafkaRouter) RegisterReplyHandler(topic string, handlerFunc KafkaReplyHandlerFunc) {
	kr.logger.Debug("subscribing to topic ", topic)
	kr.replyHandlers[topic] = handlerFunc
}

func (kr KafkaRouter) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (kr KafkaRouter) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (kr KafkaRouter) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		kr.Handle(msg.Topic, *msg)
		session.MarkMessage(msg, "")
	}
	return nil
}

func (kr KafkaRouter) Handle(topic string, message sarama.ConsumerMessage) {
	handlerFunc, ok := kr.replyHandlers[message.Topic]
	replyTopic, err := findHeader(message.Headers, "kafka_replyTopic")
	if err != nil {
		kr.logger.Error(err)
	}
	correlationID, err := findHeader(message.Headers, "kafka_correlationId")
	if err != nil {
		kr.logger.Error(err)
	}
	msg := KafkaMessage{ConsumerMessage: message, ReplyTopic: replyTopic}
	if ok {
		resp, err := handlerFunc(msg)
		if err != nil {
			kr.logger.Error(err)
		}
		if resp != nil {
			kr.kafkaClient.Send(replyTopic, resp, correlationID)
		}
	}
}
