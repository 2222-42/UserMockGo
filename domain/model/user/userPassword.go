package user

import "UserMockGo/domain/model"

type Password struct {
	ID       model.UserID
	Password PassString // TODO: add validation
}

func NewPassword(id model.UserID, password PassString) Password {
	return Password{
		ID:       id,
		Password: password,
	}
}
