package infrainterface

import (
	"UserMockGo/domain/model/authorization"
	"UserMockGo/domain/model/user"
)

type ITokenManager interface {
	GenerateToken(u user.User) (string, error)
	Parse(tokenString string) (authorization.Authorization, error)
	//RevokeToken(str string) error
}
