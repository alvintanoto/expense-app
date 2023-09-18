package service

import (
	"expense_app/internal/entity"
	"expense_app/internal/repository"
	"expense_app/internal/session"
	"expense_app/internal/util/config"
	"expense_app/internal/util/logger"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthenticationService interface {
		Login(username, password string) (accessToken, refreshToken string, err error)
	}

	implAuth struct {
		config  config.Configuration
		logger  logger.Logger
		session session.Session
		repo    repository.Holder
	}
)

func NewAuthenticationService(config config.Configuration, logger logger.Logger, session session.Session, repo repository.Holder) (AuthenticationService, error) {
	return &implAuth{config: config, logger: logger, session: session, repo: repo}, nil
}

func (i *implAuth) Login(username, password string) (accessToken, refreshToken string, err error) {
	user, err := i.repo.UserRepository.GetActiveUserByUsername(username)
	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		i.logger.Error("error compare hash and password")
		return "", "", err
	}

	accessToken, refreshToken, err = i.CreateUserSession(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (i *implAuth) CreateUserSession(user *entity.User) (accessToken, refreshToken string, err error) {
	accessTokenExpireDuration, err := time.ParseDuration(i.config.JwtAccessTokenExpireDuration)
	if err != nil {
		i.logger.Error("fail parsing access token expire duration")
		return "", "", err
	}
	accessTokenExpireTime := time.Now().Add(time.Duration(accessTokenExpireDuration))
	claims := &session.AuthenticationClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpireTime),
		},
	}

	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenByte, err := accessTokenClaims.SignedString([]byte(i.config.AppSecretKey))
	if err != nil {
		i.logger.Error("error creating access token")
		return "", "", err
	}

	refreshTokenExpireDuration, err := time.ParseDuration(i.config.JwtRefreshTokenExpireDuration)
	if err != nil {
		i.logger.Error("fail parsing refresh token expire duration")
		return "", "", err
	}
	refreshTokenExpireTime := time.Now().Add(time.Duration(refreshTokenExpireDuration))
	refreshTokenData := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshTokenExpireTime),
	}
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenData)
	refreshTokenBytes, err := refreshTokenClaims.SignedString([]byte(i.config.AppSecretKey))
	if err != nil {
		i.logger.Error("error creating access token")
		return "", "", err
	}

	accessToken = string(accessTokenByte)
	refreshToken = string(refreshTokenBytes)
	go i.session.CreateUserSession(user.ID, accessToken, refreshToken, user)
	return accessToken, refreshToken, nil
}
