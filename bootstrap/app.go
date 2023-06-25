package bootstrap

import (
	"main/pkg/postgresql"
)

type Application struct {
	Env      *Env
	Postgres postgresql.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresClient(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresConnection(app.Postgres)
}
