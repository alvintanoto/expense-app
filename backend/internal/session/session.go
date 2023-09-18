package session

import (
	"context"
	"encoding/json"
	"errors"
	"expense_app/internal/entity"
	"expense_app/internal/util/config"
	"expense_app/internal/util/logger"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

type (
	AuthenticationClaims struct {
		UserID string `json:"user_id"`
		jwt.RegisteredClaims
	}

	UserSessionData struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	Session interface {
		CreateUserSession(userID, accessToken, refreshToken string, user *entity.User)
		GetUserSession(userID, accessToken string) (userSessionData *UserSessionData, err error)
		GetUserRefreshSession(userID, accessToken string) (refreshToken string, err error)
		DeleteUserSession(userID, accessToken string)
	}

	sessionImpl struct {
		Config config.Configuration
		Logger logger.Logger

		RedisClient *redis.Client
	}
)

var (
	ErrInvalidRefreshToken = errors.New("invalid_refresh_token")
)

func NewSession(config config.Configuration, logger logger.Logger, redis *redis.Client) (Session, error) {
	return &sessionImpl{Config: config, RedisClient: redis, Logger: logger}, nil
}

func (i *sessionImpl) CreateUserSession(userID, accessToken, refreshToken string, user *entity.User) {
	// insert into redis
	accessTokenExpireDuration, err := time.ParseDuration(i.Config.JwtAccessTokenExpireDuration)
	if err != nil {
		i.Logger.Error("error parsing duration")
		return
	}

	userData := &UserSessionData{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	jsonBytes, err := json.Marshal(userData)
	if err != nil {
		i.Logger.Error("error marshaling json")
		return
	}

	key := fmt.Sprintf("user_session:%s:%s", userID, accessToken)
	_, err = i.RedisClient.Set(context.Background(), key, string(jsonBytes), accessTokenExpireDuration).Result()
	if err != nil {
		i.Logger.Error("error creating user session")
		return
	}

	refreshTokenExpireDuration, err := time.ParseDuration(i.Config.JwtRefreshTokenExpireDuration)
	if err != nil {
		i.Logger.Error("error parsing duration")
		return
	}

	refreshTokenKey := fmt.Sprintf("session_refresh:%s", accessToken)
	_, err = i.RedisClient.Set(context.Background(), refreshTokenKey, refreshToken, refreshTokenExpireDuration).Result()
	if err != nil {
		i.Logger.Error("error creating user session refresh token")
		return
	}
}

func (i *sessionImpl) GetUserSession(userID, accessToken string) (userSessionData *UserSessionData, err error) {
	key := fmt.Sprintf("user_session:%s:%s", userID, accessToken)
	result, err := i.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		i.Logger.Error("error getting user session data")
		return nil, err
	}

	err = json.Unmarshal([]byte(result), &userSessionData)
	if err != nil {
		i.Logger.Error("error unmarshaling user session data")
		return nil, err
	}

	return userSessionData, nil
}

func (i *sessionImpl) DeleteUserSession(userID, accessToken string) {
	key := fmt.Sprintf("user_session:%s:%s", userID, accessToken)
	refreshTokenKey := fmt.Sprintf("session_refresh:%s", accessToken)

	ctx := context.Background()
	i.RedisClient.Del(ctx, key, refreshTokenKey)
}

func (i *sessionImpl) GetUserRefreshSession(userID, accessToken string) (refreshToken string, err error) {
	refreshTokenKey := fmt.Sprintf("session_refresh:%s", accessToken)
	result, err := i.RedisClient.Get(context.Background(), refreshTokenKey).Result()
	if err != nil {
		i.Logger.Error("error getting session refresh token")
		return "", ErrInvalidRefreshToken
	}

	return result, nil
}
