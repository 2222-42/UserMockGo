package main

import (
	"UserMockGo/domain/service"
	"UserMockGo/infra/mysql"
	"UserMockGo/infra/randomintgenerator"
	"UserMockGo/web/handler"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	userRepository := mysql.UserRepositoryMock{}
	userIdGenerator := randomintgenerator.UserIdGeneratorMock{}
	userService := service.NewUserService(userRepository, userIdGenerator)
	userHandler := handler.NewUserHandler(userService)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", userHandler.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
