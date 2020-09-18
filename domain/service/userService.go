package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
	"time"
)

type UserService struct {
	userRepository infrainterface.IUserRepository
	idGenerator    infrainterface.IUserIdGenerator
}

func NewUserService(userRepository infrainterface.IUserRepository, idGenerator infrainterface.IUserIdGenerator) UserService {
	return UserService{
		userRepository: userRepository,
		idGenerator:    idGenerator,
	}
}

//Passwordはこの時点ではいらないかも？
func (service UserService) CreateUser(email user.Email, password user.PassString, passwordConfirmation user.PassString) error {
	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()

	user := user.NewUser(model.UserID(id), email, password, passwordConfirmation, now)

	// TODO: error handling
	return service.userRepository.Save(user)
}
