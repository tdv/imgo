package service

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	Storage
	client *redis.Client
}

func (this *redisCache) init() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if client == nil {
		return errors.New("Failed to create redis client.")
	}

	this.client = client

	return nil
}

func (this *redisCache) Put(id string, buf []byte) error {
	err := this.client.Set(id, buf, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (this *redisCache) Get(id string) ([]byte, error) {
	val, err := this.client.Get(id).Result()

	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func CreateRedisCache() (Storage, error) {
	client := redisCache{}

	if err := client.init(); err != nil {
		return nil, err
	}

	return &client, nil
}
