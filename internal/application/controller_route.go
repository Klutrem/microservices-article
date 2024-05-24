package application

import (
	"main/pkg"
	"main/pkg/kafka"
	"slices"
)

type ControllerRoutes struct {
	controller TaskController
	logger     pkg.Logger
	kafka      kafka.KafkaClient
}

func NewControllerRoutes(logger pkg.Logger, controller TaskController, kafka kafka.KafkaClient) ControllerRoutes {
	return ControllerRoutes{
		logger:     logger,
		controller: controller,
		kafka:      kafka,
	}
}

func (tr ControllerRoutes) Setup() {
	tr.logger.Info("setting up routes")

	kafka.AllTopics = slices.Concat(kafka.AllTopics, TaskTopics)

	kafkaRouter := kafka.NewKafkaRouter(tr.logger, tr.kafka)

	kafkaRouter.RegisterReplyHandler(test1, tr.controller.TestConsumeTopic)
	kafkaRouter.RegisterReplyHandler(test2, tr.controller.TestSecondTopic)
	kafkaRouter.RegisterReplyHandler(test3, tr.controller.TestThirdTopic)

	kafkaRouter.RegisterReplyHandler(test4, tr.controller.TestTopic4)
	kafkaRouter.RegisterReplyHandler(test5, tr.controller.TestTopic5)
	kafkaRouter.RegisterReplyHandler(test6, tr.controller.TestTopic6)

	kafkaRouter.RegisterReplyHandler(test7, tr.controller.TestTopic7)
	kafkaRouter.RegisterReplyHandler(test8, tr.controller.TestTopic8)
	kafkaRouter.RegisterReplyHandler(test9, tr.controller.TestTopic9)
	kafkaRouter.RegisterReplyHandler(test10, tr.controller.TestTopic10)

	tr.kafka.Consume(kafkaRouter, TaskTopics[:])

}
