package handler

import (
	"UserMockGo/domain/model/user"
	"UserMockGo/domain/service"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

type UserParam struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type ActivationParam struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (handler UserHandler) Create(c echo.Context) error {
	body := new(UserParam)
	if err := c.Bind(body); err != nil {
		fmt.Println("Request is failed: " + err.Error())
		return err
	}

	if body.Password != body.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, "Password does not match PasswordConfirmation")
	}

	if err := handler.userService.CreateUser(user.Email(body.Email), user.PassString(body.Password)); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, body)
}

func (handler UserHandler) Activate(c echo.Context) error {
	body := new(ActivationParam)
	if err := c.Bind(body); err != nil {
		fmt.Println("Request is failed: " + err.Error())
		return err
	}

	if err := handler.userService.ActivateUser(user.Email(body.Email), body.Token); err != nil {
		fmt.Println("Activation is failed: " + err.Error())
		return c.JSON(http.StatusBadRequest, "Activation is failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, body)
}
