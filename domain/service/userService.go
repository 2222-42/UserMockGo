package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/user"
	"UserMockGo/lib/valueObjects/userValues"
	"net/http"
	"time"
)

type UserService struct {
	userRepository     infrainterface.IUserRepository
	idGenerator        infrainterface.IUserIdGenerator
	tokenGenerator     infrainterface.IUserTokenGenerator
	activationNotifier infrainterface.IActivationNotifier
}

func NewUserService(
	userRepository infrainterface.IUserRepository,
	idGenerator infrainterface.IUserIdGenerator,
	tokenGenerator infrainterface.IUserTokenGenerator,
	activationNotifier infrainterface.IActivationNotifier,
) UserService {
	return UserService{
		userRepository:     userRepository,
		idGenerator:        idGenerator,
		tokenGenerator:     tokenGenerator,
		activationNotifier: activationNotifier,
	}
}

func (service UserService) CreateUser(email userValues.Email, passString userValues.PassString) error {

	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()
	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()
	userId := model.UserID(id)

	u, err := user.NewUser(userId, email, now)
	if err != nil {
		return err
	}

	p, err := user.NewPassword(userId, passString)
	if err != nil {
		return err
	}
	a := user.NewActivation(userId, token, expiresAt)

	if err := service.userRepository.CreateUserTransactional(u, p, a); err != nil {
		return err
	}

	return service.activationNotifier.SendEmail(u, a)
}

func (service UserService) ActivateUser(email userValues.Email, token string) error {
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

func (service UserService) ReissueOfActivation(email userValues.Email) error {
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

	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()
	a := user.NewActivation(u.ID, token, expiresAt)

	return service.userRepository.ReissueOfActivationTransactional(a)
}
