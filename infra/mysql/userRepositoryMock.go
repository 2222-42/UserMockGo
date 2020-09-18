package mysql

import "UserMockGo/domain/model/user"

type UserRepositoryMock struct {
}

func (repo UserRepositoryMock) CreateUser(user user.User, pass user.Password, activation user.Activation) error {
	return nil
}
