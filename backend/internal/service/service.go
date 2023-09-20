package service

import "go.uber.org/dig"

type Holder struct {
	dig.In

	AuthenticationService AuthenticationService
	UserService           UserService
	CurrencyService       CurrencyService
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuthenticationService); err != nil {
		return nil
	}

	if err := container.Provide(NewUserService); err != nil {
		return err
	}

	if err := container.Provide(NewCurrencyService); err != nil {
		return err
	}

	return nil
}
