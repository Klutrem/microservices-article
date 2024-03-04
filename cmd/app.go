package cmd

import (
	"context"
	"log"
	route "main/internal/application"
	"main/internal/config"
	"main/pkg"

	"go.uber.org/fx"
)


func Run() any {
	return func(
		route route.Routes,
		router pkg.RequestHandler,
		env config.Env,
	) {
		route.Setup()
		err := router.Gin.Run(":" + env.Port)
		if err != nil {
			log.Fatal(err)
		}
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
