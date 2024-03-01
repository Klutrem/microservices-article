package TaskRoute

import (
	"main/api/controller"
	"main/domain/domainCommon"
	"main/lib"
)

type TaskRouter struct {
	controller controller.TaskController
	handler    lib.RequestHandler
	kafka      lib.KafkaClient
	env        lib.Env
	logger     lib.Logger
}

func NewTaskRouter(controller controller.TaskController, handler lib.RequestHandler, env lib.Env, kafka lib.KafkaClient, logger lib.Logger) TaskRouter {
	return TaskRouter{
		controller: controller,
		handler:    handler,
		kafka:      kafka,
		env:        env,
		logger:     logger,
	}
}

func (tr TaskRouter) Setup() {
	group := tr.handler.Gin.Group("")
	group.GET("/kafka", tr.controller.TestReplyTopic)

	kafkaRouter := domainCommon.NewKafkaRouter(tr.logger, tr.kafka)
	kafkaRouter.RegisterReplyHandler("topic.test", tr.controller.TestConsumeTopic)

	go tr.kafka.Consume(kafkaRouter, lib.Testtopics[:])
}
