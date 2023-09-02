package util

import (
	"expense_app/internal/util/config"
	"expense_app/internal/util/logger"

	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	Config config.Configuration
	Logger logger.Logger
}

func Register(container *dig.Container) error {
	if err := container.Provide(logger.NewLogger); err != nil {
		return err
	}

	if err := container.Provide(config.New); err != nil {
		return err
	}

	return nil
}
