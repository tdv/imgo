package main

import (
	"bufio"
	"fmt"
	"imgo/service"
	"os"
)

func main() {
	if config, err := service.InitConfig(); err != nil {
		panic("Failed to read config. Error: " + err.Error())
	} else if builder, err := service.CreateAppBuilder(config); err != nil {
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
