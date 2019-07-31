package service

type redisCache struct {
	Storage
}

func (this *redisCache) Put(id string, buf []byte) error {
	return nil
}

func (this *redisCache) Get(id string) ([]byte, error) {
	return nil, nil
}

func CreateRedisCache() (Storage, error) {
	return &redisCache{}, nil
}
