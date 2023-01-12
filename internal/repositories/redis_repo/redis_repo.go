package redis_repo

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gramilul123/test-makves/internal/models"
	"github.com/gramilul123/test-makves/pkg/errors_handler"
	"github.com/gramilul123/test-makves/pkg/logger"
)

// RedisRepo -.
type RedisRepo struct {
	client *redis.Client
	logger *logger.ZapLogger
}

// New -.
func NewRedisRepo(client *redis.Client, logger *logger.ZapLogger) *RedisRepo {

	return &RedisRepo{client: client, logger: logger}
}

func (s RedisRepo) Set(ctx context.Context, in map[string]*models.User) error {

	err := s.client.MSet(ctx, in)

	if err != nil {
		s.logger.Error("error while setting users", err)

		return errors_handler.ErrInternalService
	}

	return nil
}

func (s RedisRepo) Get(ctx context.Context, in []string) ([]*models.User, error) {
	users := []*models.User{}

	res := s.client.MGet(ctx, in...)
	if res.Err() != nil {
		s.logger.Error("error while geting users", res.Err())

		return nil, errors_handler.ErrInternalService
	}

	if err := res.Scan(&users); err != nil {
		s.logger.Error("error while scan result", err)

		return nil, errors_handler.ErrInternalService
	}

	return users, nil
}
