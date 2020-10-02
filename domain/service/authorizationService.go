package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/authorization"
)

type AuthorizationService struct {
	tokenManager infrainterface.ITokenManager
}

func NewAuthorizationService(tokenManager infrainterface.ITokenManager) AuthorizationService {
	return AuthorizationService{
		tokenManager: tokenManager,
	}
}

func (service AuthorizationService) GetAuthorization(tokenString string) (authorization.Authorization, error) {
	auth, err := service.tokenManager.Parse(tokenString)

	if err != nil {
		return authorization.Authorization{}, err
	}

	return auth, err
}
