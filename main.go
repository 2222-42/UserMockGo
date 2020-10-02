package main

import (
	"UserMockGo/domain/service"
	"UserMockGo/infra/encryption"
	"UserMockGo/infra/jwtManager"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/notifier"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/infra/token"
	"UserMockGo/web/handler"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	userRepository := mysql.NewUserRepositoryMock()
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationNotifier := notifier.NewActivationNotifier()
	LoginInfra := encryption.NewLoginInfraMock()
	tokenManager := jwtManager.NewTokenManagerMock()
	userService := service.NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationNotifier, LoginInfra, tokenManager)
	userHandler := handler.NewUserHandler(userService)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", userHandler.Create)
	e.GET("/user/activate", userHandler.Activate)
	e.POST("/user/reissue", userHandler.Reissue)
	e.POST("/user/login", userHandler.Login)
	e.Logger.Fatal(e.Start(":8080"))
}
