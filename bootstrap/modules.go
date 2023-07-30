package bootstrap

import (
	"main/api/controller"
	"main/api/route"
	TaskRoute "main/api/route/task"
	UserRoute "main/api/route/user"
	"main/infrastructure"
	"main/lib"
	"main/usecase"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	infrastructure.Module,
	usecase.Module,
	route.Module,
	lib.Module,
	controller.Module,
	TaskRoute.Module,
	UserRoute.Module,
)
