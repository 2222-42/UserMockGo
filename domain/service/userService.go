package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/user"
	"net/http"
	"time"
)

type UserService struct {
	userRepository       infrainterface.IUserRepository
	idGenerator          infrainterface.IUserIdGenerator
	tokenGenerator       infrainterface.IUserTokenGenerator
	activationRepository infrainterface.IActivationRepository
}

func NewUserService(
	userRepository infrainterface.IUserRepository,
	idGenerator infrainterface.IUserIdGenerator,
	tokenGenerator infrainterface.IUserTokenGenerator,
	activationRepository infrainterface.IActivationRepository,
) UserService {
	return UserService{
		userRepository:       userRepository,
		idGenerator:          idGenerator,
		tokenGenerator:       tokenGenerator,
		activationRepository: activationRepository,
	}
}

//Passwordはこの時点ではいらないかも？
func (service UserService) CreateUser(email user.Email, passString user.PassString) error {
	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()
	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()
	userId := model.UserID(id)

	u := user.NewUser(userId, email, now)
	p := user.NewPassword(userId, passString)
	a := user.NewActivation(userId, token, expiresAt)

	// TODO: transactional commit
	return service.userRepository.CreateUserTransactional(u, p, a)
}

func (service UserService) ActivateUser(email user.Email, token string) error {
	u, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	a, err := service.activationRepository.FindByUserIdAndToken(u.ID, token)
	if err != nil {
		return err
	}

	if !a.IsValid() {
		return errors.MyError{
			StatusCode: http.StatusForbidden,
			Message:    "expired",
			ErrorType:  "activation_token_is_expired",
		}
	}

	//TODO: update User
	return nil

}
