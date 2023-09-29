package repository

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserRepository     UserRepository
	CurrencyRepository CurrencyRepository
	WalletRepository   WalletRepository
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserRepository); err != nil {
		return err
	}

	if err := container.Provide(NewCurrencyRepository); err != nil {
		return err
	}

	if err := container.Provide(NewWalletRepository); err != nil {
		return err
	}

	return nil
}
