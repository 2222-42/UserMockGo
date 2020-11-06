package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
)

type MfaService struct {
	emailNotifier infrainterface.IEmailNotifier
	mfaManager    infrainterface.IMfaManager
}

func NewMfaService(
	activationNotifier infrainterface.IEmailNotifier,
	mfaManager infrainterface.IMfaManager,
) MfaService {
	return MfaService{
		mfaManager:    mfaManager,
		emailNotifier: activationNotifier,
	}
}

func (service MfaService) SendCode(user user.User) error {
	code := service.mfaManager.GenerateCode(user)
	return service.emailNotifier.SendCode(user, code)
}
