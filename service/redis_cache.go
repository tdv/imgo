package service

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type redisCache struct {
	Storage
	client *redis.Client

	expirationTimeout time.Duration
}

func (this *redisCache) init(address string, password string, db int, expiration time.Duration) error {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	if client == nil {
		return errors.New("Failed to create redis client.")
	}

	this.expirationTimeout = expiration
	this.client = client

	return nil
}

func (this *redisCache) Put(id string, buf []byte) error {
	err := this.client.Set(id, buf, this.expirationTimeout).Err()
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

func CreateRedisCache(config *viper.Viper) (Storage, error) {
	client := redisCache{}

	if err := client.init(
		config.GetString("cache.redis.address"),
		config.GetString("cache.redis.password"),
		config.GetInt("cache.redis.db"),
		time.Duration(config.GetInt("cache.redis.expiration"))*time.Minute,
	); err != nil {
		return nil, err
	}

	return &client, nil
}
