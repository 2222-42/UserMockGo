package authorization

import (
	"UserMockGo/domain/model"
	"UserMockGo/lib/valueObjects/userValues"
)

type Authorization struct {
	UserId model.UserID
	Email  userValues.Email
}
