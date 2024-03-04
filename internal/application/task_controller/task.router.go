package task_controller

import (
	"main/internal/config"
	"main/pkg"
	"main/pkg/kafka"
	"slices"
)



type TaskRouter struct {
	controller TaskController
	handler    pkg.RequestHandler
	kafka      kafka.KafkaClient
	logger     pkg.Logger
	env config.Env
}

func NewTaskRouter(controller TaskController,
	handler pkg.RequestHandler, 
	env config.Env, 
	kafka kafka.KafkaClient, 
	logger pkg.Logger) TaskRouter {
	return TaskRouter{
		controller: controller,
		handler:    handler,
		kafka:      kafka,
		env:        env,
		logger:     logger,
	}
}

func (tr TaskRouter) Setup() {
	kafka.AllTopics = slices.Concat(kafka.AllTopics, TaskTopics)

	kafkaRouter := kafka.NewKafkaRouter(tr.logger, tr.kafka)
	kafkaRouter.RegisterReplyHandler(firstTopic, tr.controller.TestConsumeTopic)

	go tr.kafka.Consume(kafkaRouter, TaskTopics[:])
}
