package redis_client

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/gramilul123/test-makves/config"
	"github.com/gramilul123/test-makves/pkg/errors_handler"
	"github.com/gramilul123/test-makves/pkg/logger"
)

type Redis struct {
	Client *redis.Client
}

func NewClient(logger *logger.ZapLogger, config *config.Redis) (*redis.Client, error) {

	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       config.DbNumber, // use default DB
	})

	if err := rds.Ping(context.Background()).Err(); err != nil {

		//logger.Error("Error while redis init") //TODO: Исправить логгер

		return nil, errors_handler.ErrInternalService
	}

	return rds, nil
}
