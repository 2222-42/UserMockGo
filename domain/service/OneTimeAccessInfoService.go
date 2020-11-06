package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
)

type OneTimeAccessInfoService struct {
	oneTimeAccessInfoRepository infrainterface.IOneTimeAccessInfoRepository
	mfaRepository               infrainterface.IMfaManager
	tokenManager                infrainterface.ITokenManager
	userRepository              infrainterface.IUserRepository
}

func NewOneTimeAccessInfoService(
	oneTimeAccessInfoRepository infrainterface.IOneTimeAccessInfoRepository,
	mfaRepository infrainterface.IMfaManager,
	tokenManager infrainterface.ITokenManager,
	userRepository infrainterface.IUserRepository,
) OneTimeAccessInfoService {
	return OneTimeAccessInfoService{
		oneTimeAccessInfoRepository: oneTimeAccessInfoRepository,
		mfaRepository:               mfaRepository,
		tokenManager:                tokenManager,
		userRepository:              userRepository,
	}
}

func (service OneTimeAccessInfoService) Generate(userId model.UserID) string {
	code := service.oneTimeAccessInfoRepository.CreateOneTimeAccessInfo(userId)
	return code
}

func (service OneTimeAccessInfoService) CheckWithMfaAndOneTimeCode(oneTimeCode string, mfaCode string) (string, error) {
	userId, err := service.oneTimeAccessInfoRepository.GetUserIdByOneTimeCode(oneTimeCode)
	if err != nil {
		return "", err
	}

	if err := service.mfaRepository.RequireValidPairOfUserAndCode(userId, mfaCode); err != nil {
		service.oneTimeAccessInfoRepository.IncrementRetryCount(oneTimeCode)
		return "", err
	}

	service.oneTimeAccessInfoRepository.RemoveAccessInfo(oneTimeCode)

	u, err := service.userRepository.FindById(userId)
	if err != nil {
		return "", err
	}

	token, err := service.tokenManager.GenerateToken(u)
	if err != nil {
		return "", err
	}

	return token, nil
}
