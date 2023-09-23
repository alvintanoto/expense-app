package controller

import (
	"expense_app/internal/controller/dto"
	"expense_app/internal/repository"
	"expense_app/internal/service"
	"expense_app/internal/util/logger"
	"expense_app/internal/util/response"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	RegistrationHandler interface {
		RegisterUser(ctx echo.Context) error
	}

	implRegistration struct {
		logger   logger.Logger
		services service.Holder
	}
)

func NewRegistrationHandler(logger logger.Logger, services service.Holder) (RegistrationHandler, error) {
	return &implRegistration{logger: logger, services: services}, nil
}

func (i *implRegistration) RegisterUser(ctx echo.Context) error {
	var payload dto.RegisterRequestDTO
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

	if err := i.services.UserService.Register(payload.Username, payload.Email, payload.Password); err != nil {
		switch err {
		case repository.ErrorConstraintViolation:
			return ctx.JSON(response.UsernameOrEmailAlreadyExist.HttpCode, response.UsernameOrEmailAlreadyExist.Response)
		}
		return ctx.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
	}

	accessToken, refreshToken, err := i.services.AuthenticationService.Login(payload.Username, payload.Password)
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return ctx.JSON(response.IncorrectCredentialsError.HttpCode, response.IncorrectCredentialsError.Response)
		default:
			return ctx.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
		}
	}

	resp := response.Success
	resp.Response.Data = &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return ctx.JSON(resp.HttpCode, resp.Response)
}
