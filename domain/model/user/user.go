package user

type User struct {
	ID                       int64  // TODO: make value object and add type
	Email                    string // TODO: add type and make validation
	Password                 string // TODO: add validation
	PasswordConfirmation     string
	IsActive                 bool
	ActivationToken          string
	ActivationTokenExpiresAt int64
	CreatedAt                int64
	UpdatedAt                int64
}

// TODO: tokenの生成と有効期限の設定は外部に切り出す。
func NewUser(id int64, email string, password string, passwordConfirmation string, now int64) User {
	return User{
		ID:                       id,
		Email:                    email,
		Password:                 password,
		PasswordConfirmation:     passwordConfirmation,
		IsActive:                 false,
		ActivationToken:          "",
		ActivationTokenExpiresAt: now + 60*60,
		CreatedAt:                now,
		UpdatedAt:                now,
	}
}
