package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/user"
	"UserMockGo/domain/model/valueObjects"
	"net/http"
	"time"
)

type UserService struct {
	userRepository infrainterface.IUserRepository
	idGenerator    infrainterface.IUserIdGenerator
	tokenGenerator infrainterface.IUserTokenGenerator
}

func NewUserService(
	userRepository infrainterface.IUserRepository,
	idGenerator infrainterface.IUserIdGenerator,
	tokenGenerator infrainterface.IUserTokenGenerator,
) UserService {
	return UserService{
		userRepository: userRepository,
		idGenerator:    idGenerator,
		tokenGenerator: tokenGenerator,
	}
}

func (service UserService) CreateUser(email valueObjects.Email, passString user.PassString) error {

	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()
	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()
	userId := model.UserID(id)

	u, err := user.NewUser(userId, email, now)
	if err != nil {
		return err
	}

	p := user.NewPassword(userId, passString)
	a := user.NewActivation(userId, token, expiresAt)

	return service.userRepository.CreateUserTransactional(u, p, a)
}

func (service UserService) ActivateUser(email valueObjects.Email, token string) error {
	u, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	if u.IsActive {
		return errors.MyError{
			StatusCode: http.StatusForbidden,
			Message:    "The user is already activated.",
			ErrorType:  "user_not_needed_to_activate",
		}
	}

	a, err := service.userRepository.FindByUserIdAndToken(u.ID, token)
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

	if err := service.userRepository.ActivateUserTransactional(u, a); err != nil {
		return err
	}

	return nil
}
