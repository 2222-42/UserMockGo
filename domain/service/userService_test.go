package service

import (
	"UserMockGo/infra/jwtManager"
	"UserMockGo/infra/mfa"
	"UserMockGo/infra/myBcryption"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/notifier"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/infra/token"
	"testing"
)

func TestUserService_LoginSuccess(t *testing.T) {
	userRepository := mysql.NewUserRepositoryMock()
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	oneTimeAccessInfoRepo := mysql.NewOneTimeAccessInfoRepositoryMock()

	userService := NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationNotifier, loginInfra, tokenManager, mfaManager, oneTimeAccessInfoRepo)
	code, err := userService.Login("test3@test.com", "test123456")
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

func TestUserService_LoginFail(t *testing.T) {
	userRepository := mysql.NewUserRepositoryMock()
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	oneTimeAccessInfoRepo := mysql.NewOneTimeAccessInfoRepositoryMock()
	userService := NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationNotifier, loginInfra, tokenManager, mfaManager, oneTimeAccessInfoRepo)
	code, err := userService.Login("test@test.com", "testtesttesttesttest")

	if err == nil {
		t.Error("should succeed", err)
	}

	if code != "" {
		t.Error("should get code")
	}
}
