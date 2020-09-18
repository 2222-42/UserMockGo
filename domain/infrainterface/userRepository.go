package infrainterface

import "UserMockGo/domain/model/user"

type IUserRepository interface {
	CreateUser(user user.User, pass user.Password, activation user.Activation) error
}
