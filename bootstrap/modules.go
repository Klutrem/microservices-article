package bootstrap

import (
	"main/api/controller"
	"main/api/route"
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
)
