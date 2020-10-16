package main

import (
	"UserMockGo/domain/service"
	"UserMockGo/infra/bcrypt"
	"UserMockGo/infra/jwtManager"
	"UserMockGo/infra/mfa"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/notifier"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/infra/table"
	"UserMockGo/infra/token"
	"UserMockGo/lib/valueObjects/userValues"
	"UserMockGo/web/handler"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"time"
)

func main() {
	e := echo.New()
	sampleUser := table.User{
		ID:        123,
		Email:     os.Getenv("USER_MOCK_GO_USER_EMAIL"),
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	hashedPass, _ := bcrypt.HashPassString(userValues.PassString(os.Getenv("USER_MOCK_GO_USER_PASS")))
	samplePass := table.Password{
		ID:       123,
		Password: hashedPass,
	}
	userRepository := mysql.NewUserRepositoryMock(sampleUser, samplePass)
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	LoginInfra := bcrypt.NewLoginInfraMock()
	MfaManager := mfa.NewMfaManagerMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	userService := service.NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationNotifier, LoginInfra, tokenManager)
	authorizationService := service.NewAuthorizationService(tokenManager)
	mfaService := service.NewMfaService(userRepository, activationNotifier, tokenManager, MfaManager)
	userHandler := handler.NewUserHandler(userService, authorizationService)
	mfaHandler := handler.NewMfaHandler(mfaService, authorizationService)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", userHandler.Create)
	e.GET("/user/activate", userHandler.Activate)
	e.POST("/user/reissue", userHandler.Reissue)
	e.POST("/user/login", userHandler.Login)
	e.POST("/user/mfa", mfaHandler.MFAuthenticate)
	e.GET("/users", userHandler.GetUserInfo)
	e.Logger.Fatal(e.Start(":8080"))
}
