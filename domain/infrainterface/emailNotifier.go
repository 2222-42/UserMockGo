package infrainterface

import "UserMockGo/domain/model/userModel"

type IEmailNotifier interface {
	SendActivationEmail(user userModel.User, activation userModel.Activation, subjectStr string) error
	SendCode(user userModel.User, code string) error
}
