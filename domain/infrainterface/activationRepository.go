package infrainterface

import (
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
)

type IActivationRepository interface {
	FindByUserIdAndToken(userId model.UserID, token string) (user.Activation, error)
}
