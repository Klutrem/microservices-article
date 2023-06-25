package route

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewUserRouter),
	fx.Provide(NewTaskRouter),
	fx.Provide(NewRoutes),
)

type Routes []Route

// Route interface
type Route interface {
	Setup()
}

func NewRoutes(
	taskroutes TaskRouter,
	userroutes UserRouter,
) Routes {
	return Routes{
		taskroutes,
		userroutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
