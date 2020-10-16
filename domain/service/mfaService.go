package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
)

type MfaService struct {
	emailNotifier infrainterface.IEmailNotifier
	mfaManager    infrainterface.IMfaManager
	tokenManager  infrainterface.ITokenManager
}

func (service MfaService) SendCode(user user.User) error {
	code := service.mfaManager.GenerateCode(user)
	return service.emailNotifier.SendCode(user, code)
}

func (service MfaService) CheckCode(user user.User, code string) (string, error) {
	if err := service.mfaManager.RequireValidPairOfUserAndCode(user, code); err != nil {
		return "", err
	}

	token, err := service.tokenManager.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
