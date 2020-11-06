package infrainterface

import (
	"UserMockGo/domain/model/authorizationModel"
	"UserMockGo/domain/model/userModel"
)

type ITokenManager interface {
	GenerateToken(u userModel.User) (string, error)
	Parse(tokenString string) (authorizationModel.Authorization, error)
	//RevokeToken(str string) error
}
