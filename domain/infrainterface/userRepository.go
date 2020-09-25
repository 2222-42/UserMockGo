package infrainterface

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
)

type IUserRepository interface {
	CreateUserTransactional(user user.User, pass user.Password, activation user.Activation) error
	FindByEmail(email user.Email) (user.User, error)
	ActivateUserTransactional(user user.User, activation user.Activation) error
	FindByUserIdAndToken(userId model.UserID, token string) (user.Activation, error)
}
