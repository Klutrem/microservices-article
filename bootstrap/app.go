package bootstrap

import (
	"context"
	"log"
	"main/lib"

	"main/api/route"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

type Application struct {
	Command *cobra.Command
}

func Run() interface{} {
	return func(
		route route.Routes,
		router lib.RequestHandler,
		env lib.Env,
	) {
		route.Setup()
		err := router.Gin.Run(":" + env.Port)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetCobraCommand(opt fx.Option) *cobra.Command {
	Command := &cobra.Command{
		Use: "main",
		Run: func(cmd *cobra.Command, args []string) {
			opts := fx.Options(
				fx.Invoke(Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer func() {
				err := app.Stop(ctx)
				if err != nil {
					log.Fatal(err)
				}
			}()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	return Command
}

var rootCmd = &cobra.Command{
	Use:   "clean-gin",
	Short: "Clean architecture using gin framework",
}

func NewApp() Application {
	cmd := Application{
		Command: rootCmd,
	}
	cmd.Command.AddCommand(GetCobraCommand(CommonModules))
	return cmd
}

var RootApp = NewApp()
