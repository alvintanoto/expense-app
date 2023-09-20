package controller

import (
	"expense_app/internal/repository"
	"expense_app/internal/service"
	"expense_app/internal/util/logger"
	"expense_app/internal/util/response"

	"github.com/labstack/echo/v4"
)

type (
	CurrencyHandler interface {
		GetCurrencies(ctx echo.Context) error
	}

	implCurrency struct {
		logger  logger.Logger
		service service.Holder
	}
)

func NewCurrencyHandler(logger logger.Logger, service service.Holder) (CurrencyHandler, error) {
	return &implCurrency{logger: logger, service: service}, nil
}

func (i *implCurrency) GetCurrencies(ctx echo.Context) error {
	result, err := i.service.CurrencyService.GetCurrencies()
	if err != nil {
		if err == repository.ErrorRecordNotFound {
			return ctx.JSON(response.RecordNotFoundError.HttpCode, response.RecordNotFoundError.Response)
		}

		return ctx.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
	}

	resp := response.Success
	resp.Response.Data = result
	return ctx.JSON(resp.HttpCode, resp.Response)
}
