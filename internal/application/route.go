package application

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewControllerRoutes),
	fx.Provide(NewController),
)

type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// here should return routers
func NewRoutes(controllerRoutes ControllerRoutes) Routes {
	return Routes{
		controllerRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
