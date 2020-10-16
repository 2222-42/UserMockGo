package infrainterface

import "UserMockGo/domain/model/user"

type IEmailNotifier interface {
	SendActivationEmail(user user.User, activation user.Activation, subjectStr string) error
	SendCode(user user.User, code string) error
}
