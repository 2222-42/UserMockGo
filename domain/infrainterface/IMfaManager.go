package infrainterface

import "UserMockGo/domain/model/user"

type IMfaManager interface {
	GenerateCode(user user.User) string
	RequireValidPairOfUserAndCode(user user.User, code string) error
}
