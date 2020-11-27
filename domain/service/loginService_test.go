package service

import (
	"UserMockGo/infra/jwtManager"
	"UserMockGo/infra/mfa"
	"UserMockGo/infra/myBcryption"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/notifier"
	"testing"
)

func TestLoginService_LoginSuccess(t *testing.T) {
	userRepository := mysql.NewUserRepositoryMock()
	activationNotifier := notifier.NewActivationNotifierMock()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	oneTimeAccessInfoRepo := mysql.NewOneTimeAccessInfoRepositoryMock()

	loginService := NewLoginService(userRepository, loginInfra, oneTimeAccessInfoRepo, mfaManager, activationNotifier)
	code, err := loginService.Login("test3@test.com", "test123456")
	if err != nil {
		t.Error("failed", err)
	}

	if code == "" {
		t.Error("should get code")
	}

	oneTimeService := NewOneTimeAccessInfoService(oneTimeAccessInfoRepo, mfaManager, tokenManager, userRepository)
	accessToken, err := oneTimeService.CheckWithMfaAndOneTimeCode(code, "123456")
	if err != nil {
		t.Error("failed", err)
	}

	if accessToken == "" {
		t.Error("should get token")
	}
}

func TestLoginService_LoginFail(t *testing.T) {
	userRepository := mysql.NewUserRepositoryMock()
	activationNotifier := notifier.NewActivationNotifierMock()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	oneTimeAccessInfoRepo := mysql.NewOneTimeAccessInfoRepositoryMock()
	loginService := NewLoginService(userRepository, loginInfra, oneTimeAccessInfoRepo, mfaManager, activationNotifier)

	code, err := loginService.Login("test@test.com", "testtesttesttesttest")

	if err == nil {
		t.Error("should succeed", err)
	}

	if code != "" {
		t.Error("should get code")
	}
}
