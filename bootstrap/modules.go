package bootstrap

import (
	"main/api/controller"
	"main/api/route"
	TaskRoute "main/api/route/task"
	"main/lib"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	route.Module,
	lib.Module,
	controller.Module,
	TaskRoute.Module,
)
