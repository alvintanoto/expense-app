package controller

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	RegistrationHandler RegistrationHandler

	CurrencyHandler CurrencyHandler
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewRegistrationHandler); err != nil {
		return err
	}

	if err := container.Provide(NewCurrencyHandler); err != nil {
		return err
	}

	return nil
}

func (c Holder) Routes(app *echo.Echo) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/register", c.RegistrationHandler.RegisterUser)

	// TODO: user authentication & authorization

	// TODO: Cache currency data 24Hour
	v1.GET("/currencies", c.CurrencyHandler.GetCurrencies)

	// TODO: wallets
}
