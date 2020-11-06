package infrainterface

import "UserMockGo/domain/model"

type IOneTimeAccessInfoRepository interface {
	CreateOneTimeAccessInfo(userId model.UserID) string
	GetUserIdByOneTimeCode(code string) (model.UserID, error)
	RemoveAccessInfo(code string)
	IncrementRetryCount(code string)
}
