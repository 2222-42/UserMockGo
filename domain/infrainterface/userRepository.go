package infrainterface

import "UserMockGo/domain/model/user"

type IUserRepository interface {
	Save(user user.User) error
}
