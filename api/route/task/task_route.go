package TaskRoute

import (
	"main/api/controller"
	"main/lib"
)

type TaskRouter struct {
	controller   controller.TaskController
	handler      lib.RequestHandler
	kafkaHandler lib.KafkaHandler
	kafka        lib.KafkaClient
	env          lib.Env
}

func NewTaskRouter(controller controller.TaskController, handler lib.RequestHandler, env lib.Env, kafka lib.KafkaClient, kafkaHandler lib.KafkaHandler) TaskRouter {
	return TaskRouter{
		controller:   controller,
		handler:      handler,
		kafka:        kafka,
		env:          env,
		kafkaHandler: kafkaHandler,
	}
}

func (tr TaskRouter) Setup() {
	group := tr.handler.Gin.Group("")
	group.GET("/kafka", tr.controller.TestReplyTopic)

	go tr.kafka.Consume(tr.kafkaHandler, lib.Testtopics[:])
}
