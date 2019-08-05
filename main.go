package main

import (
	"bufio"
	"fmt"
	"imgo/service"
	"os"
)

func main() {
	var err error

	var stg service.Storage
	if stg, err = service.CreatePostgresStorage(); err != nil {
		panic("Failed to create storage. Error: " + err.Error())
	}

	var cache service.Storage
	if cache, err = service.CreateRedisCache(); err != nil {
		panic("Failed to create cache. Error: " + err.Error())
	}

	var conv service.Converter
	if conv, err = service.CreateImageMagickConverter(); err != nil {
		panic("Failed to create converter. Error: " + err.Error())
	}

	var srv service.Service
	if srv, err = service.CreateHttpServer(conv, stg, cache); err != nil {
		panic("Failed to start server. Error: " + err.Error())
	}
	defer srv.Stop()
	srv.Start()

	fmt.Println("ImGo has been started successfully!")
	fmt.Println("Press Enter for quit.")

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
