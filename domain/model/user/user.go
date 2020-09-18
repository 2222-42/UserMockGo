package user

import "UserMockGo/domain/model"

type Email string
type PassString string

type User struct {
	ID                       model.UserID
	Email                    Email      // TODO: add type and make validation
	Password                 PassString // TODO: add validation
	PasswordConfirmation     PassString
	IsActive                 bool
	ActivationToken          string
	ActivationTokenExpiresAt int64
	CreatedAt                int64
	UpdatedAt                int64
}

// TODO: tokenの生成と有効期限の設定は外部に切り出す。
func NewUser(id model.UserID, email Email, password PassString, passwordConfirmation PassString, now int64, token string, expiresAt int64) User {
	return User{
		ID:                       id,
		Email:                    email,
		Password:                 password,
		PasswordConfirmation:     passwordConfirmation,
		IsActive:                 false,
		ActivationToken:          token,
		ActivationTokenExpiresAt: expiresAt,
		CreatedAt:                now,
		UpdatedAt:                now,
	}
}
