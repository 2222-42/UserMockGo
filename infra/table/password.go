package table

import (
	"UserMockGo/domain/model/user"
	"UserMockGo/infra/bcrypt"
)

type Password struct {
	ID       int64  `gorm:"primary_key;not null;column:id;"`
	Password string `gorm:"not null;column:password;"`
}

func MapFromUserPasswordModel(password user.Password) (Password, error) {
	hp, err := bcrypt.HashPassString(password.Password)
	if err != nil {
		return Password{}, err
	}

	return Password{
		ID:       int64(password.ID),
		Password: hp,
	}, nil
}

func (p Password) MapToHashedString() string {
	return p.Password
}
