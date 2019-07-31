package service

type httpService struct {
	Service
}

func (this *httpService) Start() error {
	return nil
}

func (this *httpService) Stop() error {
	return nil
}

func CreateHttpServer() (Service, error) {
	return &httpService{}, nil
}
