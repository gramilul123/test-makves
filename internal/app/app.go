package app

import (
	"fmt"
	"log"

	"github.com/gramilul123/test-makves/config"
	"github.com/gramilul123/test-makves/pkg/logger"
	"github.com/gramilul123/test-makves/pkg/redis_client"
	//cv1 "github.com/gramilul123/test-makves/internal/http/controllers/v1"
	//"github.com/gramilul123/test-makves/internal/http/middlewares"
	//"github.com/gramilul123/test-makves/internal/http/routes"
	//"github.com/gramilul123/test-makves/internal/repositories/postgres_repo"
	//"github.com/gramilul123/test-makves/internal/usecases"
	//"github.com/gramilul123/test-makves/pkg/gin"
	//"github.com/gramilul123/test-makves/pkg/logger"
	//"github.com/gramilul123/test-makves/pkg/postgres"
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

	/*pg, err := postgres.New(cfg.Db)
	if err != nil {
		l.Error(err.Error(), "Postgres start error")

		return
	}
	defer pg.Close()

	// Repositories
	pgRepo := postgres_repo.NewPostrgesRepo(pg.Db, l)

	// Gin
	ginHandler := gin.NewRequestHandler(l)

	// Middlewares
	middlewares.NewMiddlewares(
		middlewares.NewTimeoutMiddleware(&ginHandler, l, cfg),
	).Setup()

	// usecases
	helpersUsecase := usecases.NewHelpersUsecase(pgRepo, l)
	themesUsecase := usecases.NewThemesUsecase(pgRepo, l)
	messagesUsecase := usecases.NewMessagesUsecase(pgRepo, l)
	usersUsecase := usecases.NewUsersUsecase(pgRepo, l)
	reportUsecase := usecases.NewReportUsecasee(pgRepo, l)
	servicesUsecase := usecases.NewServicesUsecase(pgRepo, l)

	// controllers
	helpersController := cv1.NewHelpersController(helpersUsecase, l)
	themesController := cv1.NewThemesController(themesUsecase, l)
	messagesController := cv1.NewMessagesController(messagesUsecase, l)
	usersController := cv1.NewUsersController(usersUsecase, l)
	reportController := cv1.NewReportController(reportUsecase, l)
	servicesController := cv1.NewServicesController(servicesUsecase, l)

	//routes
	helpersRoutes := routes.NewHelpersRoutes(l, ginHandler, helpersController, cfg.App)
	themesRoutes := routes.NewThemesRoutes(l, ginHandler, themesController)
	messagesRoutes := routes.NewMessagesRoutes(l, ginHandler, messagesController)
	usersRoutes := routes.NewUsersRoutes(l, ginHandler, usersController)
	reportRoutes := routes.NewReportRoutes(l, ginHandler, reportController)
	servicesRoutes := routes.NewServicesRoutes(l, ginHandler, servicesController)

	routes.NewRoutes(
		helpersRoutes,
		themesRoutes,
		messagesRoutes,
		usersRoutes,
		reportRoutes,
		servicesRoutes,
	).Setup()

	err = ginHandler.Gin.Run(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
	if err != nil {
		l.Error(err.Error(), "Gin start error")

		return
	}*/
}
