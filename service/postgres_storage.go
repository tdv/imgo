package service

type postgresStorage struct {
	Storage
}

func (this *postgresStorage) Put(id string, buf []byte) error {
	return nil
}

func (this *postgresStorage) Get(id string) ([]byte, error) {
	return nil, nil
}

func CreatePostgresStorage() (Storage, error) {
	return &postgresStorage{}, nil
}
