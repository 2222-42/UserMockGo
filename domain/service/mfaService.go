package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
)

type MfaService struct {
	emailNotifier  infrainterface.IEmailNotifier
	mfaManager     infrainterface.IMfaManager
	tokenManager   infrainterface.ITokenManager
	userRepository infrainterface.IUserRepository
}

func NewMfaService(
	userRepository infrainterface.IUserRepository,
	activationNotifier infrainterface.IEmailNotifier,
	tokenManager infrainterface.ITokenManager,
	mfaManager infrainterface.IMfaManager,
) MfaService {
	return MfaService{
		mfaManager:     mfaManager,
		userRepository: userRepository,
		emailNotifier:  activationNotifier,
		tokenManager:   tokenManager,
	}
}

func (service MfaService) SendCode(user user.User) error {
	code := service.mfaManager.GenerateCode(user)
	return service.emailNotifier.SendCode(user, code)
}

func (service MfaService) CheckCode(userId model.UserID, code string) (string, error) {
	if err := service.mfaManager.RequireValidPairOfUserAndCode(userId, code); err != nil {
		return "", err
	}

	u, err := service.userRepository.FindById(userId)
	if err != nil {
		return "", err
	}

	token, err := service.tokenManager.GenerateToken(u, true)
	if err != nil {
		return "", err
	}

	return token, nil
}
