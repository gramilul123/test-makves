package gin

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/gramilul123/test-makves/pkg/logger"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(l *logger.ZapLogger) RequestHandler {
	engine := gin.New()

	engine.Use(ginzap.Ginzap(l.Log, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(l.Log, true))

	return RequestHandler{Gin: engine}
}
