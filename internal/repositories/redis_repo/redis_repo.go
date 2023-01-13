package redis_repo

import (
	"context"
	"encoding/json"
	"fmt"

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

func (s RedisRepo) Set(ctx context.Context, in map[string]string) error {

	err := s.client.MSet(ctx, in).Err()
	if err != nil {
		s.logger.Error(fmt.Sprintf("error while set user: %s", err))

		return errors_handler.ErrInternalService
	}

	return nil
}

func (s RedisRepo) Get(ctx context.Context, in []string) ([]*models.User, error) {
	users := []string{}

	res := s.client.MGet(ctx, in...)
	if res.Err() != nil {
		s.logger.Error(fmt.Sprintf("error while geting users: %s", res.Err()))

		return nil, errors_handler.ErrInternalService
	}

	if err := res.Scan(&users); err != nil {
		s.logger.Error(fmt.Sprintf("error while scan result: %s", err))

		return nil, errors_handler.ErrInternalService
	}

	userList := make([]*models.User, len(users))
	for _, user := range users {
		err := json.Unmarshal([]byte(user), &userList)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error parse to struct: %s", err))

			return nil, errors_handler.ErrInternalService
		}
	}

	return userList, nil
}
