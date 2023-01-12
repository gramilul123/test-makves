package makves_repo

import (
	"context"
	"io"
	"net/http"

	"github.com/gocarina/gocsv"
	"github.com/gramilul123/test-makves/internal/models"
	"github.com/gramilul123/test-makves/pkg/errors_handler"
	"github.com/gramilul123/test-makves/pkg/logger"
)

// MakvesRepo -.
type MakvesRepo struct {
	httpClient *http.Client
	logger     *logger.ZapLogger
}

// New -.
func NewMakvesRepo(httpClient *http.Client, logger *logger.ZapLogger) *MakvesRepo {

	return &MakvesRepo{httpClient: httpClient, logger: logger}
}

func (s MakvesRepo) Download(ctx context.Context, url string) ([]*models.User, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		s.logger.Error("error while create request", err)

		return nil, errors_handler.ErrInternalService
	}

	resp, err := s.httpClient.Do(request)
	if err != nil {
		s.logger.Error("error while http connect", err, url)

		return nil, errors_handler.ErrHttpConnection
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Error("bad status", resp.StatusCode, url)

		return nil, errors_handler.ErrHttpConnection
	}

	csvBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("error while serialize body", err)

		return nil, errors_handler.ErrInternalService
	}

	users := []*models.User{}

	if err := gocsv.UnmarshalBytes(csvBytes, &users); err != nil {
		s.logger.Error("error while parse body", err)

		return nil, errors_handler.ErrInternalService
	}

	return users, nil
}
