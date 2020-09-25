package notifier

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
)

type activationNotifier struct {
}

func NewActivationNotifier() infrainterface.IActivationNotifier {
	return activationNotifier{}
}

func (notifier activationNotifier) SendEmail(user user.User, activation user.Activation) error {
	//email := user.Email
	//token := activation.ActivationToken

	return nil
}
