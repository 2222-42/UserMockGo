package authorization

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/lib/valueObjects/userValues"
	"net/http"
)

type Authorization struct {
	UserId model.UserID
	Email  userValues.Email
}

func (auth Authorization) RequireSameUser(id model.UserID) error {

	if auth.UserId != id {
		return errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "not same id",
			ErrorType:  "invalid_authorization",
		}
	}

	return nil
}
