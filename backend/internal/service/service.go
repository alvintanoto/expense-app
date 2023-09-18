package service

import "go.uber.org/dig"

type Holder struct {
	dig.In

	AuthenticationService AuthenticationService
	UserService           UserService
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuthenticationService); err != nil {
		return nil
	}

	if err := container.Provide(NewUserService); err != nil {
		return err
	}

	return nil
}
