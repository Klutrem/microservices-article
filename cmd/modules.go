package cmd

import (
	route "main/internal/application"
	"main/internal/application/task_controller"
	"main/internal/config"
	"main/pkg"
	"main/pkg/kafka"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	kafka.Module,
	pkg.Module,
	task_controller.Module,
	route.Module,
)
