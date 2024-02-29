package TaskRoute

import (
	"main/api/controller"
	"main/lib"

	"github.com/IBM/sarama"
)

type EventHandler struct {
	controller controller.TaskController
}

func NewKafkaHandler(controller controller.TaskController) lib.KafkaHandler {
	return EventHandler{
		controller: controller,
	}
}

func (handler EventHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (handler EventHandler) Cleanup(session sarama.ConsumerGroupSession) error {
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
	replyTopic := topic + ".reply"
	switch topic {
	case lib.TestTopic:
		handler.controller.TestConsumeTopic(replyTopic, message)
	case lib.SecondTopic:
		handler.controller.TestSecondTopic(replyTopic, message)
	}
}
