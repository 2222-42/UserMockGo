package initializer

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/infra/mysql"
)

type Repositories struct {
	userRepository        infrainterface.IUserRepository
	oneTimeAccessInfoRepo infrainterface.IOneTimeAccessInfoRepository
}

func InitRepositories() *Repositories {
	userRepository := mysql.NewUserRepositoryMock()
	oneTimeAccessInfoRepo := mysql.NewOneTimeAccessInfoRepositoryMock()
	return &Repositories{
		userRepository:        userRepository,
		oneTimeAccessInfoRepo: oneTimeAccessInfoRepo,
	}
}
