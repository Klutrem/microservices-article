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
	group := tr.handler.Gin.Group("").Use(middleware.JwtAuthMiddleware(tr.env.AccessTokenSecret))
	group.GET("/task", tr.controller.Fetch)
	group.POST("/task", tr.controller.Create)
}
