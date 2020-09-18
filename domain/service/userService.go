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
func (service UserService) CreateUser(email user.Email, passString user.PassString) error {
	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()
	userId := model.UserID(id)

	u := user.NewUser(userId, email, now)
	p := user.NewPassword(userId, passString)
	a := user.NewActivation(userId, "", now+60*60)

	// TODO: transactional commit
	return service.userRepository.CreateUser(u, p, a)
}
