package table

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
)

type Activation struct {
	ID                       int64  `gorm:"primary_key;not null;column:id;"`
	ActivationToken          string `gorm:"not null;column:activation_token;default:''"`
	ActivationTokenExpiresAt int64  `gorm:"not null; column: activation_token_expires_at; default: 0"`
}

func MapFromUserActivationModel(activation user.Activation) (Activation, error) {
	return Activation{
		ID:                       int64(activation.ID),
		ActivationToken:          activation.ActivationToken,
		ActivationTokenExpiresAt: activation.ActivationTokenExpiresAt,
	}, nil
}

func (a Activation) MapToActivationModel() (user.Activation, error) {
	return user.Activation{
		ID:                       model.UserID(a.ID),
		ActivationToken:          a.ActivationToken,
		ActivationTokenExpiresAt: a.ActivationTokenExpiresAt,
	}, nil
}
