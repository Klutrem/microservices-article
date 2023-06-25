package bootstrap

import (
	"fmt"
	"log"

	"main/pkg/postgresql"
)

func NewPostgresClient(env *Env) postgresql.Client {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	var dbURI string
	if dbUser == "" || dbPass == "" {
		dbURI = fmt.Sprintf("postgres://%s:%s/%s", dbHost, dbPort, dbName)
	} else {
		dbURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	}
	client, err := postgresql.NewClient(dbURI)
	if err != nil {
		log.Fatal(err.Error())
	}

	return client
}

func ClosePostgresConnection(client postgresql.Client) {
	if client == nil {
		return
	}

	client.Disconnect()
}
