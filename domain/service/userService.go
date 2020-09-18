package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
	"math/rand"
	"time"
)

type UserService struct {
	userRepository infrainterface.IUserRepository
}

func NewUserService(userRepository infrainterface.IUserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

//Passwordはこの時点ではいらないかも？
func (service UserService) CreateUser(email string, password string, passwordConfirmation string) error {
	// TODO: idの生成
	rand.Seed(time.Now().Unix())
	id := rand.Int63n(10000)
	// TODO: timerを導入する
	now := time.Now().Unix()

	user := user.NewUser(id, email, password, passwordConfirmation, now)

	// TODO: error handling
	return service.userRepository.Save(user)
}
