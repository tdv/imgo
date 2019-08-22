package main

import (
	"bufio"
	"fmt"
	"imgo/service"
	"os"
)

func main() {
	if builder, err := service.CreateAppBuilder(); err != nil {
		panic("Failed to create application. Error: " + err.Error())
	} else if app, err := builder.Build(); err != nil {
		panic("Failed to create server. Error: " + err.Error())
	} else {
		server := app.(service.Service)
		server.Start()
		defer server.Stop()

		fmt.Println("ImGo has been started successfully!")
		fmt.Println("Press Enter for quit.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
