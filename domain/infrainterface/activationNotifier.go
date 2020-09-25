package infrainterface

import "UserMockGo/domain/model/user"

type IActivationNotifier interface {
	SendEmail(user user.User, activation user.Activation) error
}
