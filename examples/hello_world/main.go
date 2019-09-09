// HelloWorld example shows how you can use 'ImGo'
// as a package for build your own application
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tdv/imgo/service"
)

func main() {
	if config, err := service.LoadDefConfig(); err != nil {
		panic("Failed to read config. Error: " + err.Error())
	} else if builder, err := service.InitAppBuilder(config); err != nil {
		panic("Failed to create application. Error: " + err.Error())
	} else if app, err := builder.Build(); err != nil {
		panic("Failed to create server. Error: " + err.Error())
	} else {
		server := app.(service.Service)
		server.Start()
		defer server.Stop()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		fmt.Println("ImGo has been started successfully!")
		fmt.Println("Press Ctrl+C for quit.")

		<-c
	}
}
