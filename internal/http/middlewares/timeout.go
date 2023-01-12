package middlewares

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	"github.com/gramilul123/test-makves/config"
	ginpkg "github.com/gramilul123/test-makves/pkg/gin"
	"github.com/gramilul123/test-makves/pkg/logger"
)

// TimeoutMiddleware middleware
type TimeoutMiddleware struct {
	handler *ginpkg.RequestHandler
	logger  *logger.ZapLogger
	env     *config.Config
}

func NewTimeoutMiddleware(handler *ginpkg.RequestHandler, logger *logger.ZapLogger, cfg *config.Config) TimeoutMiddleware {
	return TimeoutMiddleware{
		handler: handler,
		logger:  logger,
		env:     cfg,
	}
}

func (m TimeoutMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	m.handler.Gin.Use(timeout.New(
		timeout.WithTimeout(time.Duration(m.env.Timeout)*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
	))
}
