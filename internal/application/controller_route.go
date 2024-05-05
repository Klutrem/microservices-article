package application

import (
	"main/pkg"
)

type ControllerRoutes struct {
	controller Controller
	logger     pkg.Logger
	handler    pkg.RequestHandler
}

func NewControllerRoutes(logger pkg.Logger, handler pkg.RequestHandler, controller Controller) ControllerRoutes {
	return ControllerRoutes{
		logger:     logger,
		handler:    handler,
		controller: controller,
	}
}

func (cr ControllerRoutes) Setup() {
	cr.logger.Info("setting up routes")

	cr.handler.Gin.GET("test1", cr.controller.Test1)
	cr.handler.Gin.GET("test2", cr.controller.Test2)
	cr.handler.Gin.GET("test3", cr.controller.Test3)
}
