// Service is the images service package (ImGo)
package service

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type httpService struct {
	Service
	address   string
	converter Converter
	storage   Storage
	cache     Storage
	server    *http.Server
}

func (this *httpService) onPut(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/plain")

	var buf []byte
	var err error
	if buf, err = ioutil.ReadAll(req.Body); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Failed to process request. Error: " + err.Error()))
		return
	}

	if buf == nil || len(buf) < 1 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("No request body."))
		return
	}

	params := req.URL.Query()

	format := params.Get("format")

	if len(format) < 1 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("The \"format\" parameter must not be empty."))
		return
	}

	var width *int
	if w := params.Get("width"); len(w) > 0 {
		if val, err := strconv.Atoi(w); err != nil {
			// TODO: to log
		} else {
			width = &val
		}
	}

	var height *int
	if h := params.Get("height"); len(h) > 0 {
		if val, err := strconv.Atoi(h); err != nil {
			// TODO: to log
		} else {
			height = &val
		}
	}

	var hash string
	if buf, hash, err = this.converter.Convert(buf, format, width, height); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Failed to convert image. Error: " + err.Error()))
		return
	}

	if err = this.storage.Put(hash, buf); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Failed to store image. Error: " + err.Error()))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(hash))
}

func (this *httpService) onGet(resp http.ResponseWriter, req *http.Request) {
	items := strings.Split(req.URL.Path, "/")
	if len(items) != 3 {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Header().Add("Content-Type", "text/plain")
		resp.Write([]byte("Too long path."))
		return
	}

	id := items[2]

	var image []byte
	var err error
	if image, err = this.cache.Get(id); err != nil {
		if image, err = this.storage.Get(id); err != nil {
			resp.WriteHeader(http.StatusNotFound)
			resp.Header().Add("Content-Type", "text/plain")
			resp.Write([]byte("Image with id \"" + id + "\" not found. Message: " + err.Error()))
			return
		} else if err = this.cache.Put(id, image); err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Header().Add("Content-Type", "text/plain")
			resp.Write([]byte("Failed to put image into cache. Error: " + err.Error()))
			return
		}
	}

	resp.WriteHeader(http.StatusOK)
	resp.Header().Add("Content-Type", "image")
	resp.Write(image)
}

func (this *httpService) Start() {
	if this.server != nil {
		panic("Server already started.")
	}

	router := mux.NewRouter()

	router.HandleFunc("/put", this.onPut).Methods("PUT")
	router.HandleFunc("/get/{id}", this.onGet).Methods("GET")

	this.server = &http.Server{
		Addr:    this.address,
		Handler: router,
	}

	go func() {
		if err := this.server.ListenAndServe(); err != nil {
			panic("Failed to listen.")
		}
	}()
}

func (this *httpService) Stop() error {
	if this.server == nil {
		return errors.New("The server is not running.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	this.server.Shutdown(ctx)
	this.server = nil

	return nil
}

func (this *httpService) Started() bool {
	return this.server != nil
}

// ImplHttp - id of default http server implementation of Service interface
const ImplHttp = "http"

var _ = RegisterEntity(
	EntityServer,
	ImplHttp,
	func(ctx BuildContext) (interface{}, error) {
		config := ctx.GetConfig()
		converter := ctx.GetEntity(EntityImageConverter).(Converter)
		storage := ctx.GetEntity(EntityStorage).(Storage)
		cache := ctx.GetEntity(EntityCache).(Storage)

		server := httpService{
			address:   config.GetStrVal("address"),
			converter: converter,
			storage:   storage,
			cache:     cache,
		}

		return &server, nil
	},
)
