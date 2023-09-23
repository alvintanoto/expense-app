package controller

import (
	"expense_app/internal/session"
	"expense_app/internal/util/config"
	"expense_app/internal/util/logger"
	"expense_app/internal/util/response"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type (
	Middleware interface {
		LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc
		CustomHTTPErrorHandler(err error, c echo.Context)

		IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc
		RefreshTokenAuthentication(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewareImpl struct {
		logger  logger.Logger
		config  config.Configuration
		session session.Session
	}
)

func NewMiddleware(logger logger.Logger, config config.Configuration, session session.Session) (Middleware, error) {
	return &middlewareImpl{logger: logger, config: config, session: session}, nil
}

func (i *middlewareImpl) CustomHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		switch he.Code {
		case http.StatusNotFound:
			c.JSON(response.RouteNotFoundError.HttpCode, response.RouteNotFoundError.Response)
		default:
			c.JSON(response.InternalServerError.HttpCode, response.InternalServerError.Response)
		}
	}
}

func (i *middlewareImpl) LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		req := c.Request()
		res := c.Response()

		fields := map[string]interface{}{
			"remote_ip":  c.RealIP(),
			"latency":    time.Since(start).String(),
			"host":       req.Host,
			"method":     req.Method,
			"uri":        req.RequestURI,
			"status":     res.Status,
			"size":       res.Size,
			"user_agent": req.UserAgent(),
		}

		n := res.Status
		switch {
		case n >= 500:
			i.logger.ErrorWithCtx(c.Request().Context(), "server error", logger.ToField("data", fields))
		case n >= 400:
			i.logger.WarnWithCtx(c.Request().Context(), "client error", logger.ToField("data", fields))
		case n >= 300:
			i.logger.InfoWithCtx(c.Request().Context(), "redirection", logger.ToField("data", fields))
		default:
			i.logger.InfoWithCtx(c.Request().Context(), "success", logger.ToField("data", fields))
		}

		return nil
	}
}

func (i *middlewareImpl) IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		auth := ctx.Request().Header.Get("authorization")
		if auth == "" {
			return ctx.JSON(response.InvalidTokenError.HttpCode, response.InvalidTokenError.Response)
		}

		tokenString := strings.Split(auth, " ")[1]
		claims := &session.AuthenticationClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(i.config.AppSecretKey), nil
		})
		if err != nil {
			i.logger.Errorf("error: %s", err.Error())
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				return ctx.JSON(response.TokenExpiredError.HttpCode, response.TokenExpiredError.Response)
			case jwt.ValidationErrorSignatureInvalid:
				return ctx.JSON(response.TokenSignatureInvalid.HttpCode, response.TokenSignatureInvalid.Response)
			default:
				return ctx.JSON(response.TokenExpiredError.HttpCode, response.TokenExpiredError.Response)
			}
		}

		// check to redis
		sessionData, err := i.session.GetUserSession(claims.UserID, tokenString)
		if err != nil {
			i.logger.Error("could not validate user session token")
			return ctx.JSON(response.InvalidTokenError.HttpCode, response.InvalidTokenError.Response)
		}

		ctx.Set("session", sessionData)
		ctx.Set("claims", claims)
		ctx.Set("token", tokenString)
		return next(ctx)
	}
}

func (i *middlewareImpl) RefreshTokenAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		auth := ctx.Request().Header.Get("authorization")
		if auth == "" {
			return ctx.JSON(response.InvalidTokenError.HttpCode, response.InvalidTokenError.Response)
		}

		tokenString := strings.Split(auth, " ")[1]
		claims := &session.AuthenticationClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(i.config.AppSecretKey), nil
		})
		if err != nil {
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				return ctx.JSON(response.TokenSignatureInvalid.HttpCode, response.TokenSignatureInvalid.Response)
			}
		}

		ctx.Set("claims", claims)
		ctx.Set("token", tokenString)
		return next(ctx)
	}
}
