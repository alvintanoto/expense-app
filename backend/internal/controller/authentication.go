package controller

import (
	"expense_app/internal/controller/dto"
	"expense_app/internal/repository"
	"expense_app/internal/service"
	"expense_app/internal/session"
	"expense_app/internal/util/logger"
	"expense_app/internal/util/response"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthenticationHandler interface {
		Login(ctx echo.Context) (err error)
		CheckToken(ctx echo.Context) (err error)
		RefreshToken(ctx echo.Context) (err error)
		Logout(ctx echo.Context) (err error)
	}

	implAuthentication struct {
		logger   logger.Logger
		services service.Holder
	}
)

func NewAuthenticationHandler(logger logger.Logger, services service.Holder) (AuthenticationHandler, error) {
	return &implAuthentication{logger: logger, services: services}, nil
}

func (i *implAuthentication) Login(ctx echo.Context) (err error) {
	var payload dto.LoginRequestDTO

	err = ctx.Bind(&payload)
	if err != nil {
		return ctx.JSON(response.BadRequestError.HttpCode, response.BadRequestError.Response)
	}

	if err := payload.Validate(); err != nil {
		i.logger.Error(err)
		badRequestResp := response.BadRequestError
		badRequestResp.Response.ClientMessage = err.Error()
		return ctx.JSON(response.BadRequestError.HttpCode, response.BadRequestError.Response)
	}

	accessToken, refreshToken, err := i.services.AuthenticationService.Login(payload.Username, payload.Password)
	if err != nil {
		switch err {
		case repository.ErrorRecordNotFound:
			return ctx.JSON(response.UserNotFoundError.HttpCode, response.UserNotFoundError.Response)
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

func (i *implAuthentication) CheckToken(ctx echo.Context) (err error) {
	userSession := ctx.Get("session").(*session.UserSessionData)

	resp := response.Success
	resp.Response.Data = userSession
	return ctx.JSON(resp.HttpCode, resp.Response)
}

func (i *implAuthentication) RefreshToken(ctx echo.Context) (err error) {
	var (
		payload     dto.RefreshTokenDTO
		claims      *session.AuthenticationClaims
		accessToken string
	)

	claims = ctx.Get("claims").(*session.AuthenticationClaims)
	accessToken = ctx.Get("token").(string)

	if err = ctx.Bind(&payload); err != nil {
		i.logger.Error("error binding data")
		return ctx.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
	}

	accessToken, refreshToken, err := i.services.AuthenticationService.RefreshToken(claims.UserID, accessToken, payload.RefreshToken)
	if err != nil {
		switch err {
		case session.ErrInvalidRefreshToken:
			return ctx.JSON(response.RefreshTokenInvalid.HttpCode, response.RefreshTokenInvalid.Response)
		case repository.ErrorRecordNotFound:
			return ctx.JSON(response.UserNotFoundError.HttpCode, response.UserNotFoundError.Response)
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

func (i *implAuthentication) Logout(ctx echo.Context) (err error) {
	var (
		claims *session.AuthenticationClaims
		token  string
	)

	claims = ctx.Get("claims").(*session.AuthenticationClaims)
	token = ctx.Get("token").(string)

	go i.services.AuthenticationService.Logout(claims.UserID, token)

	resp := response.Success
	resp.Response.Data = nil
	return ctx.JSON(resp.HttpCode, resp.Response)
}
