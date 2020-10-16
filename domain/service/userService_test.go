package service

import (
	"UserMockGo/infra/jwtManager"
	"UserMockGo/infra/mfa"
	"UserMockGo/infra/myBcryption"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/notifier"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/infra/table"
	"UserMockGo/infra/token"
	"UserMockGo/lib/valueObjects/userValues"
	"testing"
	"time"
)

func TestUserService_LoginSuccess(t *testing.T) {
	sampleUser := table.User{
		ID:        123,
		Email:     "test@test.com",
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	hashedPass, _ := myBcryption.HashPassString(userValues.PassString("testtesttesttest"))
	samplePass := table.Password{
		ID:       123,
		Password: hashedPass,
	}
	userRepository := mysql.NewUserRepositoryMock(sampleUser, samplePass)
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	userService := NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationNotifier, loginInfra, tokenManager, mfaManager)
	code, err := userService.Login("test@test.com", "testtesttesttest")
	if err != nil {
		t.Error("failed", err)
	}

	if code == "" {
		t.Error("should get code")
	}

	//TODO: 出力されたコードについては、特定のユーザー情報のアクセスのためのJWTの発行のための情報になっていることが期待される。
}

func TestUserService_LoginFail(t *testing.T) {
	sampleUser := table.User{
		ID:        123,
		Email:     "test@test.com",
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	hashedPass, _ := myBcryption.HashPassString(userValues.PassString("testtesttesttest"))
	samplePass := table.Password{
		ID:       123,
		Password: hashedPass,
	}
	userRepository := mysql.NewUserRepositoryMock(sampleUser, samplePass)
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	loginInfra := myBcryption.NewLoginInfraMock()
	mfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	userService := NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationNotifier, loginInfra, tokenManager, mfaManager)
	code, err := userService.Login("test@test.com", "testtesttesttesttest")

	if err == nil {
		t.Error("should succeed", err)
	}

	if code != "" {
		t.Error("should get code")
	}
}