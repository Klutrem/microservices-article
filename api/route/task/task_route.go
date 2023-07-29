package TaskRoute

import (
	"main/api/controller"
	"main/api/middleware"
	"main/lib"
)

type TaskRouter struct {
	controller controller.TaskController
	handler    lib.RequestHandler
	kafka      lib.KafkaClient
	env        lib.Env
}

func NewTaskRouter(controller controller.TaskController, handler lib.RequestHandler, env lib.Env, kafka lib.KafkaClient) TaskRouter {
	return TaskRouter{
		controller: controller,
		handler:    handler,
		kafka:      kafka,
		env:        env,
	}
}

func (tr TaskRouter) Setup() {
	group := tr.handler.Gin.Group("").Use(middleware.JwtAuthMiddleware(tr.env.PublicKey))
	group.GET("/task", tr.controller.Fetch)
	group.POST("/task", tr.controller.Create)

	kafkaHandler := NewKafkaHandler()

	go tr.kafka.Consume(kafkaHandler, lib.Testtopics[:])
}
