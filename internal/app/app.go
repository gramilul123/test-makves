package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gramilul123/test-makves/config"
	cv "github.com/gramilul123/test-makves/internal/http/controllers"
	"github.com/gramilul123/test-makves/internal/http/middlewares"
	"github.com/gramilul123/test-makves/internal/http/routes"
	"github.com/gramilul123/test-makves/internal/repositories/makves_repo"
	"github.com/gramilul123/test-makves/internal/repositories/redis_repo"
	"github.com/gramilul123/test-makves/internal/usecases"
	"github.com/gramilul123/test-makves/pkg/gin"
	"github.com/gramilul123/test-makves/pkg/logger"
	"github.com/gramilul123/test-makves/pkg/redis_client"
)

func Run(cfg *config.Config) {
	l, err := logger.NewZapLogger(cfg)
	if err != nil {
		log.Fatalf("error init logger: %v", err)
	}
	defer func() {
		if err := l.Sync(); err != nil {
			log.Printf("error while sync logger: %v", err)
		}
	}()

	// Redis
	l.Info(fmt.Sprintf("redis has started at %s:%s", cfg.Redis.Host, cfg.Redis.Port))
	redisClient, err := redis_client.NewClient(l, cfg.Redis)
	if err != nil {
		l.Error("failed to redis init")

		return
	}
	defer redisClient.Close()

	// Repositories
	redisRepo := redis_repo.NewRedisRepo(redisClient, l)
	makvesRepo := makves_repo.NewMakvesRepo(&http.Client{}, l)

	// Gin
	ginHandler := gin.NewRequestHandler(l)

	// Middlewares
	middlewares.NewMiddlewares(
		middlewares.NewTimeoutMiddleware(&ginHandler, l, cfg),
	).Setup()

	// usecases
	itemsUsecase := usecases.NewServiceUC(redisRepo, makvesRepo)

	// controllers
	itemsController := cv.NewItemsController(itemsUsecase, l, cfg)

	//routes
	itemsRoutes := routes.NewItemsRoutes(l, ginHandler, itemsController, cfg.App)

	routes.NewRoutes(
		itemsRoutes,
	).Setup()

	err = ginHandler.Gin.Run(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
	if err != nil {
		l.Error(err.Error(), "Gin start error")

		return
	}
}
