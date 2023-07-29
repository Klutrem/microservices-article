package TaskRoute

import (
	"main/api/controller"
	"main/lib"

	"github.com/IBM/sarama"
)

type EventHandler struct {
	controller controller.TaskController
}

func (handler EventHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (handler EventHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (handler EventHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		handler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}
	return nil
}

func (handler EventHandler) Handle(topic string, message []byte) {
	switch topic {
	case lib.TestTopic:
		handler.controller.TestConsumeTopic(message)
	case lib.SecondTopic:
		handler.controller.TestReplyTopic(message)
	}
}

func NewKafkaHandler() lib.KafkaHandler {
	return EventHandler{}
}
