package redis_repo

import (
	"context"
	"encoding/json"
	"errors"
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
	userList := make([]*models.User, 0, len(in))

	for _, key := range in {
		res := s.client.Get(ctx, key)
		if res.Err() != nil {
			if errors.Is(res.Err(), redis.Nil) {
				continue
			}

			s.logger.Error(fmt.Sprintf("error while geting users: %s", res.Err()))

			return nil, errors_handler.ErrInternalService
		}

		var userString string

		if err := res.Scan(&userString); err != nil {
			s.logger.Error(fmt.Sprintf("error while scan result: %s", err))

			return nil, errors_handler.ErrInternalService
		}

		var user models.User
		err := json.Unmarshal([]byte(userString), &user)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error parse to struct: %s", err))

			return nil, errors_handler.ErrInternalService
		}

		userList = append(userList, &user)
	}

	return userList, nil
}
