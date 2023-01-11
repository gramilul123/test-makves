package redis_client

import (
	"context"
	"fmt"

	"github.com/docker/docker/daemon/logger"
	"github.com/go-redis/redis/v8"
	"github.com/gramilul123/test-makves/config"
	"github.com/gramilul123/test-makves/pkg/errors_handler"
)

type Redis struct {
	Client *redis.Client
}

func NewClient(logger *logger.Logger, config *config.Redis) (*redis.Client, error) {

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
