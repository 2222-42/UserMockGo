package mfa

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/user"
	"net/http"
)

type MFAManagerMock struct {
}

func NewMfaManagerMock() infrainterface.IMfaManager {
	return MFAManagerMock{}
}

const testCode = "123456"

func (manager MFAManagerMock) GenerateCode(user user.User) string {
	return testCode
}

func (manager MFAManagerMock) RequireValidPairOfUserAndCode(userId model.UserID, code string) error {
	userCode, err := getCode(userId)
	if err != nil {
		return errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "check your email, password and code",
			ErrorType:  "invalid_login_info",
		}
	}

	if code != userCode {
		return errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "check your email, password and code",
			ErrorType:  "invalid_login_info",
		}
	}
	return nil
}

func getCode(userId model.UserID) (string, error) {
	if int(userId) == 0 {
		return "", errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "check your email, password and code",
			ErrorType:  "invalid_login_info",
		}
	}
	return testCode, nil
}
