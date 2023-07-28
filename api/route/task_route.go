package route

import (
	"main/api/controller"
	"main/api/middleware"
	"main/lib"
)

type TaskRouter struct {
	controller controller.TaskController
	handler    lib.RequestHandler
	env        lib.Env
}

func NewTaskRouter(controller controller.TaskController, handler lib.RequestHandler, env lib.Env) TaskRouter {
	return TaskRouter{
		controller: controller,
		handler:    handler,
		env:        env,
	}
}

func (tr TaskRouter) Setup() {
	group := tr.handler.Gin.Group("").Use(middleware.JwtAuthMiddleware(tr.env.PublicKey))
	group.GET("/task", tr.controller.Fetch)
	group.POST("/task", tr.controller.Create)

	go tr.controller.TestConsumeTopic("test")     //run all functions with kafka topics handling
	go tr.controller.TestReplyTopic("test.reply") //subscribe to the reply topic
}
