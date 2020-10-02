package handler

import (
	"UserMockGo/domain/service"
	"UserMockGo/lib/valueObjects/userValues"
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

type ReissueParam struct {
	Email string `json:"email"`
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

	if err := handler.userService.CreateUser(userValues.Email(body.Email), userValues.PassString(body.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, body)
}

func (handler UserHandler) Activate(c echo.Context) error {
	body := new(ActivationParam)
	if err := c.Bind(body); err != nil {
		fmt.Println("Request is failed: " + err.Error())
		return err
	}

	if err := handler.userService.ActivateUser(userValues.Email(body.Email), body.Token); err != nil {
		fmt.Println("Activation is failed: " + err.Error())
		return c.JSON(http.StatusBadRequest, "Activation is failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, body)
}

func (handler UserHandler) Reissue(c echo.Context) error {
	body := new(ReissueParam)
	if err := c.Bind(body); err != nil {
		fmt.Println("Request is failed: " + err.Error())
		return err
	}

	if err := handler.userService.ReissueOfActivation(userValues.Email(body.Email)); err != nil {
		fmt.Println("Reissue is failed: " + err.Error())
		return c.JSON(http.StatusBadRequest, "Reissue is failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, body)
}

func (handler UserHandler) Login(c echo.Context) error {
	body := new(LoginParams)
	if err := c.Bind(body); err != nil {
		fmt.Println("Request is failed: " + err.Error())
		return err
	}

	token, err := handler.userService.Login(userValues.Email(body.Email), userValues.PassString(body.Password))
	if err != nil {
		fmt.Println("Login is failed: " + err.Error())
		return c.JSON(http.StatusBadRequest, "Login is failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
