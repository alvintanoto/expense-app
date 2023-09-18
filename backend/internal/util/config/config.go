package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Configuration struct {
		AppSecretKey string `envconfig:"APPLICATION_SECRET_KEY" required:"true"`
		Port         string `envconfig:"PORT" default:":8080"`

		LogStdout       bool   `envconfig:"LOG_STDOUT" default:"true"`
		LogLevel        int8   `envconfig:"LOG_LEVEL" default:"-1"`
		LogFileLocation string `envconfig:"LOG_FILE_LOCATION" default:"log/application"`
		LogFileMaxAge   int    `envconfig:"LOG_FILE_MAX_AGE" default:"3"`

		DatabaseHost        string `envconfig:"DATABASE_HOST" required:"true" default:""`
		DatabaseUser        string `envconfig:"DATABASE_USER" required:"true" default:""`
		DatabasePassword    string `envconfig:"DATABASE_PASSWORD" required:"true" default:""`
		DatabaseName        string `envconfig:"DATABASE_NAME" required:"true" default:""`
		DatabasePort        string `envconfig:"DATABASE_PORT" required:"true" default:""`
		MaxOpenConns        int    `envconfig:"DATABASE_MAX_OPEN_CONNS" required:"true" default:"25"`
		MaxIdleConns        int    `envconfig:"DATABASE_MAX_IDLE_CONNS" required:"true" default:"25"`
		MaxIdleTimeDuration string `envconfig:"DATABASE_MAX_IDLE_TIME_DURATION" required:"true" default:"15m"`

		RedisDsn string `envconfig:"REDIS_DSN" required:"true" default:""`

		JwtAccessTokenExpireDuration  string `envconfig:"JWT_ACCESS_TOKEN_EXPIRE_DURATION" required:"true" default:"30m"`
		JwtRefreshTokenExpireDuration string `envconfig:"JWT_REFRESH_TOKEN_EXPIRE_DURATION" required:"true" default:"24h"`
	}
)

func New() (c Configuration, err error) {
	if err := ParseConfig(&c); err != nil {
		fmt.Println("parse config error: ", err)
		return c, err
	}

	return c, nil
}

func ParseConfig(target interface{}) error {
	var (
		filename = os.Getenv("CONFIG_FILE")
	)

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", target); err != nil {
			return err
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return err
	}

	if err := envconfig.Process("", target); err != nil {
		return err
	}

	return nil
}
