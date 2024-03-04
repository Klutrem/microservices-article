package task_controller

import (
	"main/internal/config"
	"main/pkg"
	"main/pkg/kafka"
	"slices"
)



type TaskRouter struct {
	controller TaskController
	kafka      kafka.KafkaClient
	logger     pkg.Logger
	env config.Env
}

func NewTaskRouter(controller TaskController,
	env config.Env, 
	kafka kafka.KafkaClient, 
	logger pkg.Logger) TaskRouter {
	return TaskRouter{
		controller: controller,
		kafka:      kafka,
		env:        env,
		logger:     logger,
	}
}

func (tr TaskRouter) Setup() {
	kafka.AllTopics = slices.Concat(kafka.AllTopics, TaskTopics)

	kafkaRouter := kafka.NewKafkaRouter(tr.logger, tr.kafka)
	kafkaRouter.RegisterReplyHandler(firstTopic, tr.controller.TestConsumeTopic)

	tr.kafka.Consume(kafkaRouter, TaskTopics[:])
}
