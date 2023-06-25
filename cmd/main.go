package main

import (
	"main/bootstrap"
)

func main() {
	// fx.New()

	// app := bootstrap.App()

	// env := app.Env

	// db := app.Postgres.Database()
	// defer app.CloseDBConnection()

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	// gin := gin.Default()

	// route.Setup(env, timeout, db, gin)

	// gin.Run(env.ServerAddress)
	// err := bootstrap.RootApp.Execute()
	// if err != nil {
	// 	return
	// }

	err := bootstrap.RootApp.Command.Execute()
	if err != nil {
		return
	}
}
