package routes

import (
	"github.com/gramilul123/test-makves/config"
	cv "github.com/gramilul123/test-makves/internal/http/controllers"
	"github.com/gramilul123/test-makves/pkg/gin"
	"github.com/gramilul123/test-makves/pkg/logger"
)

type ItemsRoutes struct {
	logger          *logger.ZapLogger
	handler         gin.RequestHandler
	itemsController cv.ItemsController
	config          *config.AppConfig
}

func (s ItemsRoutes) Setup() {
	s.logger.Info("Setting up routes for helpers")
	s.handler.Gin.
		GET("/set", s.itemsController.Set).
		GET("/get-items/:ids", s.itemsController.Get)
}

func NewItemsRoutes(
	logger *logger.ZapLogger,
	handler gin.RequestHandler,
	itemsController cv.ItemsController,
	config *config.AppConfig,

) ItemsRoutes {
	return ItemsRoutes{
		handler:         handler,
		logger:          logger,
		itemsController: itemsController,
		config:          config,
	}
}
