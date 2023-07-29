package main

import (
	"main/bootstrap"
)

func main() {
	//temporary for development, should be removed
	// client := gocloak.NewClient("http://localhost:8080")
	// token, err := client.LoginAdmin("admin", "admin", "aura")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(token.AccessToken)

	err := bootstrap.RootApp.Command.Execute()
	if err != nil {
		return
	}
}
