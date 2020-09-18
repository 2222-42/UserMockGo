package handler

import (
	"UserMockGo/domain/service"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (handler UserHandler) Create(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	passwordConfirmation := c.QueryParam("password_confirmation")

	return handler.userService.CreateUser(email, password, passwordConfirmation)
}
