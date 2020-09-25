package user

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/lib/valueObjects/userValues"
	"net/http"
	"time"
)

type User struct {
	ID                   model.UserID
	Email                userValues.Email
	PasswordConfirmation userValues.PassString
	IsActive             bool
	CreatedAt            int64
	UpdatedAt            int64
}

type Activation struct {
	ID                       model.UserID
	ActivationToken          string
	ActivationTokenExpiresAt int64
}

func UserNotFound(msg string) errors.MyError {
	return errors.MyError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
		ErrorType:  "user_not_found",
	}
}

func ActivationNotFound(msg string) errors.MyError {
	return errors.MyError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
		ErrorType:  "activation_not_found",
	}
}

func NewUser(id model.UserID, email userValues.Email, now int64) (User, error) {
	if !email.IsValidForm() {
		return User{}, errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "Email Form is not Valid",
			ErrorType:  "email_format_is_violated",
		}
	}

	return User{
		ID:        id,
		Email:     email,
		IsActive:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func NewActivation(id model.UserID, token string, expiresAt int64) Activation {
	return Activation{
		ID:                       id,
		ActivationToken:          token,
		ActivationTokenExpiresAt: expiresAt,
	}
}

func (a Activation) IsValid() bool {
	return a.ActivationTokenExpiresAt > time.Now().Unix()
}
