package route

import (
	"main/api/controller"
	"main/api/middleware"
	"main/lib"
)

type UserRouter struct {
	controller controller.UserController
	handler    lib.RequestHandler
	env        lib.Env
}

func NewUserRouter(controller controller.UserController, handler lib.RequestHandler, env lib.Env) UserRouter {
	return UserRouter{
		controller: controller,
		handler:    handler,
		env:        env,
	}
}

func (ur UserRouter) Setup() {
	group := ur.handler.Gin.Group("")
	group.POST("/signup", ur.controller.Signup)
	group.POST("/login", ur.controller.Login)

	protectedGroup := ur.handler.Gin.Group("").Use(middleware.JwtAuthMiddleware(ur.env.PublicKey))
	protectedGroup.GET("/profile", ur.controller.GetUserId)
}
