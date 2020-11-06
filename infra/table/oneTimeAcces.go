package table

import "UserMockGo/domain/model"

type OneTimeAccessInfo struct {
	OneTimeAccessCode string
	UserId            model.UserID
	ExpiresAt         int64
	RetryCount        int64
}
