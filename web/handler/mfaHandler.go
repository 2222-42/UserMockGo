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
	oneTimeAccessService service.OneTimeAccessInfoService
}

func NewMfaHandler(
	mfaService service.MfaService,
	authorizationService service.AuthorizationService,
	oneTimeAccessService service.OneTimeAccessInfoService,
) MfaHandler {
	return MfaHandler{
		mfaService:           mfaService,
		authorizationService: authorizationService,
		oneTimeAccessService: oneTimeAccessService,
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

	token, err := handler.oneTimeAccessService.CheckWithMfaAndOneTimeCode(c.Request().Header.Get("X-One-Time-Token"), body.Code)
	if err != nil {
		fmt.Println("Login is failed: " + err.Error())
		return c.JSON(http.StatusBadRequest, "Login is failed: "+err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
