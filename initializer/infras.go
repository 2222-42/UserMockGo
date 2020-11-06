package initializer

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/infra/jwtManager"
	"UserMockGo/infra/mfa"
	"UserMockGo/infra/myBcryption"
	"UserMockGo/infra/notifier"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/infra/token"
)

type Infras struct {
	userIdGenerator    infrainterface.IUserIdGenerator
	userTokenGenerator infrainterface.IUserTokenGenerator
	activationNotifier infrainterface.IEmailNotifier
	loginInfra         infrainterface.ILogin
	mfaManager         infrainterface.IMfaManager
	tokenManager       infrainterface.ITokenManager
}

func InitInfras() *Infras {
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	return &Infras{
		userIdGenerator:    userIdGenerator,
		userTokenGenerator: userTokenGenerator,
		activationNotifier: activationNotifier,
		loginInfra:         loginInfra,
		mfaManager:         mfaManager,
		tokenManager:       tokenManager,
	}
}
