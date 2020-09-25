package mysql

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
	"UserMockGo/domain/model/valueObjects"
	"strconv"
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

func (repo UserRepositoryMock) FindByEmail(email valueObjects.Email) (user.User, error) {
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

func (repo UserRepositoryMock) FindByUserIdAndToken(userId model.UserID, token string) (user.Activation, error) {
	if token != "" {
		switch userId {
		case 1:
			return user.Activation{
				ID:                       1,
				ActivationToken:          "aaa",
				ActivationTokenExpiresAt: 2145884400,
			}, nil
		case 2:
			return user.Activation{
				ID:                       2,
				ActivationToken:          "bbb",
				ActivationTokenExpiresAt: 0,
			}, nil
		default:
			return user.Activation{}, user.ActivationNotFound(strconv.Itoa(int(userId)))
		}
	}
	return user.Activation{}, user.ActivationNotFound(strconv.Itoa(int(userId)))
}
