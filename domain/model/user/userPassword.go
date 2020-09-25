package user

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/valueObjects"
	"net/http"
)

type Password struct {
	ID       model.UserID
	Password valueObjects.PassString
}

func NewPassword(id model.UserID, password valueObjects.PassString) (Password, error) {
	if !password.IsValidLength() {
		return Password{}, errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "The length of Password is not enough",
			ErrorType:  "password_has_not_enough_length",
		}
	}

	return Password{
		ID:       id,
		Password: password,
	}, nil
}
