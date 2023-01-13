package usecases

import (
	"context"
	"encoding/json"

	"github.com/gramilul123/test-makves/internal/models"
	"github.com/gramilul123/test-makves/internal/repositories"
	"github.com/gramilul123/test-makves/pkg/errors_handler"
	"github.com/gramilul123/test-makves/pkg/logger"
)

type Service interface {
	Set(ctx context.Context, url string) error
	Get(ctx context.Context, in []string) ([]*models.User, error)
}

type ServiceUC struct {
	logger *logger.ZapLogger
	redis  repositories.Redis
	makves repositories.Makves
}

func NewServiceUC(
	logger *logger.ZapLogger,
	redis repositories.Redis,
	makves repositories.Makves,
) *ServiceUC {

	return &ServiceUC{
		logger: logger,
		redis:  redis,
		makves: makves,
	}
}

func (f *ServiceUC) Set(ctx context.Context, url string) error {

	usersList, err := f.makves.Download(ctx, url)
	if err != nil {

		return err
	}

	usersMap := make(map[string]string, len(usersList))
	for _, user := range usersList {
		json, err := json.Marshal(user)
		if err != nil {
			f.logger.Error("error parse to json", err)

			return errors_handler.ErrInternalService
		}
		usersMap[user.Id] = string(json)
	}

	return f.redis.Set(ctx, usersMap)
}

func (f *ServiceUC) Get(ctx context.Context, in []string) ([]*models.User, error) {

	if len(in) == 0 {
		return nil, nil
	}

	return f.redis.Get(ctx, in)
}
