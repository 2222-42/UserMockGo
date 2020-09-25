package infrainterface

import "UserMockGo/domain/model/user"

type IUserRepository interface {
	CreateUserTransactional(user user.User, pass user.Password, activation user.Activation) error
	FindByEmail(email user.Email) user.User
}
