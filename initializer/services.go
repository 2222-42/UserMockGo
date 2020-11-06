package initializer

import (
	"UserMockGo/domain/service"
)

type Services struct {
	UserService              service.UserService
	OneTImeAccessInfoService service.OneTimeAccessInfoService
	AuthorizationService     service.AuthorizationService
	MfaService               service.MfaService
}

func InitServices(repositories *Repositories, infras *Infras) *Services {
	userService := service.NewUserService(repositories.userRepository, infras.userIdGenerator, infras.userTokenGenerator,
		infras.activationNotifier, infras.loginInfra, infras.tokenManager,
		infras.mfaManager, repositories.oneTimeAccessInfoRepo)
	oneTImeAccessInfoService := service.NewOneTimeAccessInfoService(repositories.oneTimeAccessInfoRepo, infras.mfaManager,
		infras.tokenManager, repositories.userRepository)
	authorizationService := service.NewAuthorizationService(infras.tokenManager)
	mfaService := service.NewMfaService(infras.activationNotifier, infras.mfaManager)
	return &Services{
		UserService:              userService,
		OneTImeAccessInfoService: oneTImeAccessInfoService,
		AuthorizationService:     authorizationService,
		MfaService:               mfaService,
	}
}
