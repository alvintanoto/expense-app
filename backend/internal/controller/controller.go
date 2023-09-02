package controller

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
}

func Register(container *dig.Container) error {
	return nil
}

func (c Holder) Routes(app *echo.Echo) {

}
