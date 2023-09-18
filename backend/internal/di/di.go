package di

import (
	"expense_app/internal/controller"
	"expense_app/internal/database"
	"expense_app/internal/repository"
	"expense_app/internal/service"
	"expense_app/internal/session"
	"expense_app/internal/util"

	"go.uber.org/dig"
)

var (
	Container = dig.New()
)

func init() {
	if err := util.Register(Container); err != nil {
		panic(err)
	}

	if err := database.Register(Container); err != nil {
		panic(err)
	}

	if err := Container.Provide(session.NewSession); err != nil {
		panic(err)
	}

	if err := controller.Register(Container); err != nil {
		panic(err)
	}

	if err := service.Register(Container); err != nil {
		panic(err)
	}

	if err := repository.Register(Container); err != nil {
		panic(err)
	}
}
