package mysql

import (
	"UserMockGo/domain/model/user"
	"time"
)

type UserRepositoryMock struct {
}

func (repo UserRepositoryMock) Save(user user.User) error {
	return nil
}

func (repo UserRepositoryMock) FindByEmail(email user.Email) user.User {
	if email == "test1@test.com" {
		return user.User{
			ID:                       1,
			Email:                    "test1@test.com",
			IsActive:                 false,
			ActivationToken:          "a",
			ActivationTokenExpiresAt: time.Now().Unix() + 60*30,
			CreatedAt:                time.Now().Unix(),
			UpdatedAt:                time.Now().Unix(),
		}
	} else if email == "test2@test.com" {
		return user.User{
			ID:                       2,
			Email:                    "test1@test.com",
			IsActive:                 false,
			ActivationToken:          "b",
			ActivationTokenExpiresAt: time.Now().Unix() - 60*10,
			CreatedAt:                time.Now().Unix() - 60*30,
			UpdatedAt:                time.Now().Unix() - 60*30,
		}
	}
	return user.User{}
}
