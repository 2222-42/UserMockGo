package infrainterface

import "UserMockGo/domain/model"

type IOneTimeAccessInfoRepository interface {
	CreateOneTimeAccessInfo(userId model.UserID, code string) error
	CheckWithCode(userId model.UserID, code string) error
}
