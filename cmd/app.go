package cmd

import (
	"context"
	route "main/internal/application"
	"main/internal/config"

	"go.uber.org/fx"
)


func Run() any {
	return func(
		route route.Routes,
		env config.Env,
	) {
		route.Setup()
	}
}


func StartApp() error {
	opts := fx.Options(
		fx.Invoke(Run()),
	)
	ctx := context.Background()
	app := fx.New(CommonModules, opts)
	err := app.Start(ctx)
	return err
}
