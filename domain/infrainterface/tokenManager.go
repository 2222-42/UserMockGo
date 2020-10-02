package infrainterface

import "UserMockGo/domain/model/user"

type ITokenManager interface {
	GenerateToken(u user.User) (string, error)
	//RevokeToken(str string) error
}
