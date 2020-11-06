package initializer

import (
	"UserMockGo/web/handler"
)

type Handlers struct {
	UserHandler handler.UserHandler
	MfaHandler  handler.MfaHandler
}

func InitHandlers(services *Services) *Handlers {
	userHandler := handler.NewUserHandler(services.UserService, services.AuthorizationService)
	mfaHandler := handler.NewMfaHandler(services.MfaService, services.AuthorizationService, services.OneTImeAccessInfoService)
	return &Handlers{
		UserHandler: userHandler,
		MfaHandler:  mfaHandler,
	}
}
