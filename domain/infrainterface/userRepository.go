package infrainterface

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/userModel"
	"UserMockGo/lib/valueObjects/userValues"
)

type IUserRepository interface {
	CreateUserTransactional(user userModel.User, pass userModel.Password, activation userModel.Activation) error
	FindByEmail(email userValues.Email) (userModel.User, error)
	FindById(id model.UserID) (userModel.User, error)
	ActivateUserTransactional(user userModel.User, activation userModel.Activation) error
	FindByUserIdAndToken(userId model.UserID, token string) (userModel.Activation, error)
	ReissueOfActivationTransactional(activation userModel.Activation) error
	GetHashedPassword(id model.UserID) (string, error)
}
