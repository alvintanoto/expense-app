package controller

import (
	"expense_app/internal/controller/dto"
	"expense_app/internal/service"
	"expense_app/internal/session"
	"expense_app/internal/util/logger"
	"expense_app/internal/util/response"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface {
		CreateUserWallet(ctx echo.Context) error // user wallet
		GetUserWallet(ctx echo.Context) error    // user walelt
	}

	implUser struct {
		logger   logger.Logger
		services service.Holder
	}
)

func NewUserHandler(logger logger.Logger, services service.Holder) (UserHandler, error) {
	return &implUser{logger: logger, services: services}, nil
}

func (i *implUser) CreateUserWallet(ctx echo.Context) error {
	var claims *session.UserSessionData
	var payload dto.UserWalletRequestDTO

	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.JSON(response.BadRequestError.HttpCode, response.BadRequestError.Response)
	}

	if err := payload.Validate(); err != nil {
		i.logger.Error(err)
		badRequestResp := response.BadRequestError
		badRequestResp.Response.ClientMessage = err.Error()
		return ctx.JSON(response.BadRequestError.HttpCode, response.BadRequestError.Response)
	}

	claims = ctx.Get("session").(*session.UserSessionData)

	if err := i.services.UserService.CreateUserWallet(claims.ID, payload.WalletName, payload.CurrencyID, payload.InitialBalance); err != nil {
		return ctx.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
	}

	resp := response.SuccessCreateUserWallet
	return ctx.JSON(resp.HttpCode, resp.Response)
}

func (i *implUser) GetUserWallet(ctx echo.Context) error {
	claims := ctx.Get("session").(*session.UserSessionData)

	wallets, err := i.services.UserService.GetUserWallet(ctx.Request().Context(), claims.ID)
	if err != nil {
		return ctx.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
	}

	resp := response.Success
	resp.Response.Data = wallets
	return ctx.JSON(resp.HttpCode, resp.Response)
}
