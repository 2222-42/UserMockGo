package infrainterface

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
)

type IMfaManager interface {
	GenerateCode(user user.User) string
	RequireValidPairOfUserAndCode(userId model.UserID, code string) error
}
