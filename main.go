package main

import (
	"UserMockGo/domain/service"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/infra/token"
	"UserMockGo/web/handler"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	userRepository := mysql.UserRepositoryMock{}
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userTokenGenerator := token.UserTokenGeneratorMock{}
	activationRepository := mysql.ActivationRepositoryMock{}
	userService := service.NewUserService(userRepository, userIdGenerator, userTokenGenerator, activationRepository)
	userHandler := handler.NewUserHandler(userService)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", userHandler.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
