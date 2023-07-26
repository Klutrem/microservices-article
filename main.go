package main

import (
	"fmt"
	"main/bootstrap"
	"main/lib"
)

func main() {
	//temporary for development, should be removed
	// client := gocloak.NewClient("http://localhost:8080")
	// token, err := client.LoginAdmin("admin", "admin", "aura")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(token.AccessToken)

	kafkaClient := lib.NewKafkaClient(lib.Env{
		BrokerHost: "localhost",
		BrokerPort: "9092",
	})

	for i := 0; i < 100; i++ {
		kafkaClient.Consume(fmt.Sprint("teeeest", i))
	}

	err := bootstrap.RootApp.Command.Execute()
	if err != nil {
		return
	}
}
