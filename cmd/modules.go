package cmd

import (
	"main/internal/application"
	"main/internal/config"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	application.Module,
)
