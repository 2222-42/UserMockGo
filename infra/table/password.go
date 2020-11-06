package table

import (
	"UserMockGo/domain/model/user"
	"UserMockGo/infra/myBcryption"
)

type Password struct {
	ID       int64  `gorm:"primary_key;not null;column:id;"`
	Password string `gorm:"not null;column:password;"`
}

func MapFromUserPasswordModel(password user.Password) (Password, error) {
	hp, err := myBcryption.HashPassString(password.Password)
	if err != nil {
		return Password{}, err
	}

	return Password{
		ID:       password.ID.ConvertUserIdToInt64(),
		Password: hp,
	}, nil
}

func (p Password) MapToHashedString() string {
	return p.Password
}
