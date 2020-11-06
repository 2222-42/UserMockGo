package table

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/userModel"
	"UserMockGo/lib/valueObjects/userValues"
)

type User struct {
	ID        int64  `gorm:"primary_key;not null;column:id;"`
	Email     string `gorm:"not null;column:email;"`
	IsActive  bool   `gorm:"not null;column:is_active;default: false"`
	CreatedAt int64  `gorm:"not null;column:created_at;default: 0"`
	UpdatedAt int64  `gorm:"not null;column:updated_at;default: 0"`
}

func MapFromUserModel(user userModel.User) (User, error) {
	return User{
		ID:        user.ID.ConvertUserIdToInt64(),
		Email:     string(user.Email),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u User) MapToUserModel() (userModel.User, error) {
	return userModel.User{
		ID:        model.UserID(u.ID),
		Email:     userValues.Email(u.Email),
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
