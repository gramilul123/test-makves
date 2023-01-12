package usecases

import (
	"context"

	"github.com/gramilul123/test-makves/internal/models"
	"github.com/gramilul123/test-makves/internal/repositories"
)

type Service interface {
	Set(ctx context.Context, url string) error
	Get(ctx context.Context, in []string) ([]*models.User, error)
}

type ServiceUC struct {
	redis  repositories.Redis
	makves repositories.Makves
}

func NewServiceUC(
	redis repositories.Redis,
	makves repositories.Makves,
) *ServiceUC {

	return &ServiceUC{
		redis:  redis,
		makves: makves,
	}
}

func (f *ServiceUC) Set(ctx context.Context, url string) error {

	usersList, err := f.makves.Download(ctx, url)
	if err != nil {

		return err
	}

	usersMap := make(map[string]*models.User, len(usersList))
	for _, user := range usersList {
		usersMap[user.Id] = user
	}

	return f.redis.Set(ctx, usersMap)
}

func (f *ServiceUC) Get(ctx context.Context, in []string) ([]*models.User, error) {

	if len(in) == 0 {
		return nil, nil
	}

	return f.redis.Get(ctx, in)
}
