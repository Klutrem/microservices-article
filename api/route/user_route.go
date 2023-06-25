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
	}
}

func (ur UserRouter) Setup() {
	group := ur.handler.Gin.Group("")
	group.POST("/signup", ur.controller.Signup)
	group.POST("/login", ur.controller.Login)
	group.POST("/refresh", ur.controller.RefreshToken)

	protected_group := ur.handler.Gin.Group("").Use(middleware.JwtAuthMiddleware(ur.env.AccessTokenSecret))
	protected_group.GET("/profile", ur.controller.Fetch)
}
