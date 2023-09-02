package database

import (
	"context"
	"database/sql"
	"expense_app/internal/util/config"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewDB); err != nil {
		return err
	}

	if err := container.Provide(NewRedisClient); err != nil {
		return err
	}

	return nil
}

func NewDB(cfg config.Configuration, logger *zap.Logger) (*sql.DB, error) {
	logger.Info("initializing database...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DatabaseHost,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabasePort,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTimeDuration)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("error pinging database")
		return nil, err
	}

	logger.Info("database initialized...")
	return db, nil
}

func NewRedisClient(cfg *config.Configuration, logger *zap.Logger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisDsn,
		Password: "",
		DB:       0,
	})

	result, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.With(zap.Error(err)).Error("error pinging redis")
		return nil, err
	}

	logger.Info("ping redis result", zap.String("result", result))
	return rdb, nil
}
