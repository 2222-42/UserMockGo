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
	tokenGenerator infrainterface.IUserTokenGenerator
}

func NewUserService(userRepository infrainterface.IUserRepository, idGenerator infrainterface.IUserIdGenerator, tokenGenerator infrainterface.IUserTokenGenerator) UserService {
	return UserService{
		userRepository: userRepository,
		idGenerator:    idGenerator,
		tokenGenerator: tokenGenerator,
	}
}

//Passwordはこの時点ではいらないかも？
func (service UserService) CreateUser(email user.Email, password user.PassString, passwordConfirmation user.PassString) error {
	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()
	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()

	u := user.NewUser(model.UserID(id), email, password, passwordConfirmation, now, token, expiresAt)

	// TODO: error handling
	return service.userRepository.Save(u)
}
