package infrainterface

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/userModel"
)

type IMfaManager interface {
	GenerateCode(user userModel.User) string
	RequireValidPairOfUserAndCode(userId model.UserID, code string) error
}
