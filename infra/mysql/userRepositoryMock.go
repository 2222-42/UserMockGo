package mysql

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
	"time"
)

type UserRepositoryMock struct {
}

func NewUserRepositoryMock() infrainterface.IUserRepository {
	return UserRepositoryMock{}
}

func (repo UserRepositoryMock) CreateUserTransactional(user user.User, pass user.Password, activation user.Activation) error {
	return nil
}

func (repo UserRepositoryMock) ActivateUserTransactional(user user.User, activation user.Activation) error {
	return nil
}

func (repo UserRepositoryMock) FindByEmail(email user.Email) (user.User, error) {
	switch email {
	case "test1@test.com":
		return user.User{
			ID:        1,
			Email:     "test1@test.com",
			IsActive:  false,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}, nil
	case "test2@test.com":
		return user.User{
			ID:        2,
			Email:     "test1@test.com",
			IsActive:  false,
			CreatedAt: time.Now().Unix() - 60*30,
			UpdatedAt: time.Now().Unix() - 60*30,
		}, nil
	default:
		return user.User{}, user.UserNotFound(string(email))
	}
}
