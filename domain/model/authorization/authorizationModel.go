package authorization

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/lib/valueObjects/userValues"
	"net/http"
)

type Authorization struct {
	UserId             model.UserID
	Email              userValues.Email
	IsMfaAuthenticated bool
}

func (auth Authorization) RequireIsMfaAuthenticated() error {
	if auth.IsMfaAuthenticated {
		return nil
	}

	return errors.MyError{
		StatusCode: http.StatusBadRequest,
		Message:    "not mfa-authenticated",
		ErrorType:  "invalid_authorization",
	}
}

func (auth Authorization) RequireSameUser(id model.UserID) error {
	if err := auth.RequireIsMfaAuthenticated(); err != nil {
		return err
	}

	if auth.UserId != id {
		return errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "not same id",
			ErrorType:  "invalid_authorization",
		}
	}

	return nil
}
