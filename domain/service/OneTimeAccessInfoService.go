package service

type OneTimeAccessInfoService struct {
}

func (service OneTimeAccessInfoService) Generate() error {
	return nil
}

func (service OneTimeAccessInfoService) CheckWithCode(code string) error {
	return nil
}
