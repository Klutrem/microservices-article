package main

import (
	"context"
	"fmt"
	"main/bootstrap"

	"github.com/Nerzal/gocloak/v13"
)

func main() {
	//temporary for development, will removed
	client := gocloak.NewClient("http://localhost:8080")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "admin", "admin", "aura")
	if err != nil {
		panic(err)
	}
	fmt.Println(token.AccessToken)

	err = bootstrap.RootApp.Command.Execute()
	if err != nil {
		return
	}

}
