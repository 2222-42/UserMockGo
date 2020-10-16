package handler

import (
	"UserMockGo/domain/service"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type MfaHandler struct {
	mfaService           service.MfaService
	authorizationService service.AuthorizationService
}

func NewMfaHandler(
	mfaService service.MfaService,
	authorizationService service.AuthorizationService,
) MfaHandler {
	return MfaHandler{
		mfaService:           mfaService,
		authorizationService: authorizationService,
	}
}

type MFAParams struct {
	Code string `json:"code"`
}

func (handler MfaHandler) MFAuthenticate(c echo.Context) error {
	body := new(MFAParams)
	if err := c.Bind(body); err != nil {
		fmt.Println("Request is failed: " + err.Error())
		return err
	}

	authorization, err := handler.authorizationService.GetAuthorization(c.Request().Header.Get("X-Access-Token"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed: "+err.Error())
	}

	token, err := handler.mfaService.CheckCode(authorization.UserId, body.Code)
	if err != nil {
		fmt.Println("Login is failed: " + err.Error())
		return c.JSON(http.StatusBadRequest, "Login is failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
