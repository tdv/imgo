package service

import (
	"errors"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

type memcachedCache struct {
	Storage
	client     *memcache.Client
	expiration int32
}

func (this *memcachedCache) init(expiration int, serverNodes ...string) error {
	client := memcache.New(serverNodes...)

	if client == nil {
		return errors.New("Failed to create memcached client.")
	}

	this.expiration = int32(expiration)
	this.client = client

	return nil
}

func (this *memcachedCache) Put(id string, buf []byte) error {
	err := this.client.Set(&memcache.Item{Key: id, Value: buf, Expiration: this.expiration})

	if err != nil {
		return err
	}

	return nil
}

func (this *memcachedCache) Get(id string) ([]byte, error) {
	val, err := this.client.Get(id)

	if err != nil {
		return nil, err
	}

	return val.Value, nil
}

const ImplMemcached = "memcached"

var _ = RegisterEntity(
	EntityCache,
	ImplMemcached,
	func(ctx BuildContext) (interface{}, error) {
		config := ctx.GetConfig()

		client := memcachedCache{}

		// The nodes separator is ';'
		// You can use some nodes which must be separated by ';'
		nodes := strings.Split(config.GetStrVal("nodes"), ";")
		if len(nodes) < 1 {
			return nil, errors.New("No nodes to connect.")
		}

		if err := client.init(
			config.GetIntVal("expiration")*60,
			nodes...,
		); err != nil {
			return nil, err
		}

		return &client, nil
	},
)
