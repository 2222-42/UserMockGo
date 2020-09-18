package mysql

import "UserMockGo/domain/model/user"

type UserRepositoryMock struct {
}

func (repo UserRepositoryMock) Save(user user.User) error {
	return nil
}
