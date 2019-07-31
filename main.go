package main

import (
	"fmt"
	"imgo/service"
)

func main() {
	var srv service.Service
	var err error
	if srv, err = service.CreateHttpServer(); err != nil {
		panic("Failed to start server. Error: " + err.Error())
	}
	defer srv.Stop()
	srv.Start()

	var stg service.Storage
	if stg, err = service.CreatePostgresStorage(); err != nil {
		panic("Failed to create storage. Error: " + err.Error())
	}
	stg.Get("") // STUB

	var cache service.Storage
	if cache, err = service.CreateRedisCache(); err != nil {
		panic("Failed to create cache. Error: " + err.Error())
	}
	cache.Get("") // STUB

	var conv service.Converter
	if conv, err = service.CreateImageMagickConverter(); err != nil {
		panic("Failed to create converter. Error: " + err.Error())
	}
	conv.Convert(make([]byte, 5, 5)) // STUB

	fmt.Println("ImGo has been started successfully!")
}
