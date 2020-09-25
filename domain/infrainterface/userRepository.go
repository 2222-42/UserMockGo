package infrainterface

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
	"UserMockGo/lib/valueObjects/userValues"
)

type IUserRepository interface {
	CreateUserTransactional(user user.User, pass user.Password, activation user.Activation) error
	FindByEmail(email userValues.Email) (user.User, error)
	ActivateUserTransactional(user user.User, activation user.Activation) error
	FindByUserIdAndToken(userId model.UserID, token string) (user.Activation, error)
	ReissueOfActivationTransactional(activation user.Activation) error
	GetHashedPassword(id model.UserID) (string, error)
}
