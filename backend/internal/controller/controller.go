package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	Middleware Middleware

	AuthenticationHandler AuthenticationHandler
	RegistrationHandler   RegistrationHandler

	CurrencyHandler CurrencyHandler
	UserHandler     UserHandler
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewMiddleware); err != nil {
		return err
	}

	if err := container.Provide(NewAuthenticationHandler); err != nil {
		return nil
	}

	if err := container.Provide(NewRegistrationHandler); err != nil {
		return err
	}

	if err := container.Provide(NewCurrencyHandler); err != nil {
		return err
	}

	if err := container.Provide(NewUserHandler); err != nil {
		return err
	}

	return nil
}

func (c Holder) Routes(app *echo.Echo) {
	c.setupEcho(app)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/register", c.RegistrationHandler.RegisterUser)

	// TODO: user authentication & authorization
	auth := v1.Group("/authentication")
	auth.POST("/login", c.AuthenticationHandler.Login)
	auth.GET("/check_token", c.AuthenticationHandler.CheckToken, c.Middleware.IsAuthenticated)
	auth.POST("/refresh_token", c.AuthenticationHandler.RefreshToken, c.Middleware.RefreshTokenAuthentication)
	auth.GET("/logout", c.AuthenticationHandler.Logout, c.Middleware.IsAuthenticated)

	// TODO: Cache currency data 24Hour
	v1.GET("/currencies", c.CurrencyHandler.GetCurrencies)

	// TODO: users (everything about user will be here)
	user := v1.Group("/user")
	user.Use(c.Middleware.IsAuthenticated)

	userWallet := user.Group("/wallet")
	userWallet.POST("", c.UserHandler.CreateUserWallet)
}

func (h *Holder) setupEcho(app *echo.Echo) {
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	app.Use(h.Middleware.LoggerMiddleware)
	app.HTTPErrorHandler = h.Middleware.CustomHTTPErrorHandler
}
