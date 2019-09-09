package main

import (
	"errors"
	"sync"

	"github.com/tdv/imgo/service"
)

type memoryCache struct {
	service.Storage
	lock     sync.Mutex
	maxItems int
	items    map[string][]byte
}

func (this *memoryCache) Put(id string, buf []byte) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	if len(this.items) >= this.maxItems {
		this.items = make(map[string][]byte)
	}
	this.items[id] = buf
	return nil
}

func (this *memoryCache) Get(id string) ([]byte, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if val, ok := this.items[id]; ok {
		return val, nil
	}
	return nil, errors.New("Item with id \"" + id + "\" not found.")
}

const ImplMemoryCache = "memory"

var _ = service.RegisterEntity(
	service.EntityCache,
	ImplMemoryCache,
	func(ctx service.BuildContext) (interface{}, error) {
		config := ctx.GetConfig()
		return &memoryCache{
			maxItems: config.GetIntVal("max_items"),
			lock:     sync.Mutex{},
			items:    make(map[string][]byte),
		}, nil
	},
)
