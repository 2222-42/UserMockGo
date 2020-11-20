package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/errors"
	"UserMockGo/lib/valueObjects/userValues"
	"net/http"
)

type LoginService struct {
	userRepository       infrainterface.IUserRepository
	loginInfra           infrainterface.ILogin
	oneTimeAccessManager infrainterface.IOneTimeAccessInfoRepository
	mfaManager           infrainterface.IMfaManager
	emailNotifier        infrainterface.IEmailNotifier
}

func NewLoginService(
	userRepository infrainterface.IUserRepository,
	loginInfra infrainterface.ILogin,
	oneTimeAccessManager infrainterface.IOneTimeAccessInfoRepository,
	mfaManager infrainterface.IMfaManager,
	emailNotifier infrainterface.IEmailNotifier,
) LoginService {
	return LoginService{
		userRepository:       userRepository,
		loginInfra:           loginInfra,
		oneTimeAccessManager: oneTimeAccessManager,
		mfaManager:           mfaManager,
		emailNotifier:        emailNotifier,
	}
}

func (service LoginService) Login(email userValues.Email, passString userValues.PassString) (string, error) {

	u, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return "", err //notValidLoginInfoError()
	}

	if !u.IsActive {
		return "", errors.MyError{
			StatusCode: http.StatusForbidden,
			Message:    "Should Authorize",
			ErrorType:  "not_valid_user_info",
		}
	}

	hp, err := service.userRepository.GetHashedPassword(u.ID)

	if err != nil {
		return "", err //notValidLoginInfoError()
	}

	if !service.loginInfra.CheckPassAndHash(hp, passString) {
		return "", notValidLoginInfoError()
	}

	token := service.oneTimeAccessManager.CreateOneTimeAccessInfo(u.ID)

	code := service.mfaManager.GenerateCode(u)
	if err := service.emailNotifier.SendCode(u, code); err != nil {
		return "", err
	}

	return token, nil
}
