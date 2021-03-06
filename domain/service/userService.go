package service

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/authorizationModel"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/userModel"
	"UserMockGo/lib/valueObjects/userValues"
	"net/http"
	"time"
)

type UserService struct {
	userRepository infrainterface.IUserRepository
	idGenerator    infrainterface.IUserIdGenerator
	tokenGenerator infrainterface.IUserTokenGenerator
	emailNotifier  infrainterface.IEmailNotifier
	tokenManager   infrainterface.ITokenManager
}

func NewUserService(
	userRepository infrainterface.IUserRepository,
	idGenerator infrainterface.IUserIdGenerator,
	tokenGenerator infrainterface.IUserTokenGenerator,
	activationNotifier infrainterface.IEmailNotifier,
	tokenManager infrainterface.ITokenManager,
) UserService {
	return UserService{
		userRepository: userRepository,
		idGenerator:    idGenerator,
		tokenGenerator: tokenGenerator,
		emailNotifier:  activationNotifier,
		tokenManager:   tokenManager,
	}
}

func notValidLoginInfoError() error {
	return errors.MyError{
		StatusCode: http.StatusForbidden,
		Message:    "Check Login Info",
		ErrorType:  "not_valid_login_info",
	}
}

func (service UserService) CreateUser(email userValues.Email, passString userValues.PassString) error {

	id := service.idGenerator.Generate()
	// TODO: timerを導入する
	now := time.Now().Unix()
	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()
	userId := model.UserID(id)

	u, err := userModel.NewUser(userId, email, now)
	if err != nil {
		return err
	}

	p, err := userModel.NewPassword(userId, passString)
	if err != nil {
		return err
	}
	a := userModel.NewActivation(userId, token, expiresAt)

	if err := service.userRepository.CreateUserTransactional(u, p, a); err != nil {
		return err
	}

	return service.emailNotifier.SendActivationEmail(u, a, "activation Account")
}

func (service UserService) ActivateUser(email userValues.Email, token string) error {
	u, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	if u.IsActive {
		return errors.MyError{
			StatusCode: http.StatusForbidden,
			Message:    "The userModel is already activated.",
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
			Message:    "The userModel is already activated.",
			ErrorType:  "user_not_needed_to_activate",
		}
	}

	token, expiresAt := service.tokenGenerator.GenerateTokenAndExpiresAt()
	a := userModel.NewActivation(u.ID, token, expiresAt)

	if err := service.userRepository.ReissueOfActivationTransactional(a); err != nil {
		return err
	}
	return service.emailNotifier.SendActivationEmail(u, a, "activation Account")
}

func (service UserService) GetUserInfo(userId model.UserID, auth authorizationModel.Authorization) (userModel.User, error) {

	if err := auth.RequireSameUser(userId); err != nil {
		return userModel.User{}, err
	}

	if auth.UserId != userId {
		return userModel.User{}, errors.MyError{
			StatusCode: http.StatusForbidden,
			Message:    "invalid user_id",
			ErrorType:  "not_accessible_this_resource",
		}
	}

	u, err := service.userRepository.FindByEmail(auth.Email)
	if err != nil {
		return userModel.User{}, err
	}

	return u, nil
}
