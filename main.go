package main

import "main/cmd"

func main() {

	err := cmd.RootApp.Command.Execute()
	if err != nil {
		return
	}
}
