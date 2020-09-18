package user

import "UserMockGo/domain/model"

type Email string
type PassString string

type User struct {
	ID                   model.UserID
	Email                Email // TODO: add type and make validation
	PasswordConfirmation PassString
	IsActive             bool
	CreatedAt            int64
	UpdatedAt            int64
}

type Activation struct {
	ID                       model.UserID
	ActivationToken          string
	ActivationTokenExpiresAt int64
}

// TODO: tokenの生成と有効期限の設定は外部に切り出す。
func NewUser(id model.UserID, email Email, now int64) User {
	return User{
		ID:        id,
		Email:     email,
		IsActive:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewActivation(id model.UserID, token string, expiresAt int64) Activation {
	return Activation{
		ID:                       id,
		ActivationToken:          token,
		ActivationTokenExpiresAt: expiresAt,
	}
}
