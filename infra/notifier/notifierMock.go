package notifier

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/userModel"
)

type notifierMock struct{}

func NewActivationNotifierMock() infrainterface.IEmailNotifier {
	return notifierMock{}
}

func (notifier notifierMock) SendActivationEmail(user userModel.User, activation userModel.Activation, subjectStr string) error {

	return nil
}

func (notifier notifierMock) SendCode(user userModel.User, code string) error {
	return nil
}
