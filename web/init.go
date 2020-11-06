package web

import (
	"UserMockGo/config"
	"UserMockGo/initializer"
	"github.com/labstack/echo"
	"net/http"
)

func Init(config *config.Config) {
	repositories := initializer.InitRepositories()
	infras := initializer.InitInfras(config)
	services := initializer.InitServices(repositories, infras)
	handlers := initializer.InitHandlers(services)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", handlers.UserHandler.Create)
	e.GET("/userModel/activate", handlers.UserHandler.Activate)
	e.POST("/userModel/reissue", handlers.UserHandler.Reissue)
	e.POST("/userModel/login", handlers.UserHandler.Login)
	e.POST("/userModel/mfa", handlers.MfaHandler.MFAuthenticate)
	e.GET("/users", handlers.UserHandler.GetUserInfo)
	e.Logger.Fatal(e.Start(":8080"))
}
