package route

import (
	TaskRoute "main/api/route/task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Routes []Route

// Route interface
type Route interface {
	Setup()
}

func NewRoutes(
	taskRoutes TaskRoute.TaskRouter,
) Routes {
	return Routes{
		taskRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
