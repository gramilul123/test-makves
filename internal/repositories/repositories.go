package repositories

import (
	"context"

	"github.com/gramilul123/test-makves/internal/models"
)

// Redis
type Redis interface {
	Set(ctx context.Context, in map[string]*models.User) error
	Get(ctx context.Context, in []string) ([]*models.User, error)
}

// Makves
type Makves interface {
	Download(ctx context.Context, url string) ([]*models.User, error)
}
