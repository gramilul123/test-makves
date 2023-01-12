package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/gramilul123/test-makves/config"
	"github.com/gramilul123/test-makves/internal/usecases"
	"github.com/gramilul123/test-makves/pkg/logger"
)

type ItemsController struct {
	usecases usecases.Service
	logger   *logger.ZapLogger
	config   *config.Config
}

func NewItemsController(
	usecases usecases.Service,
	logger *logger.ZapLogger,
	config *config.Config,
) ItemsController {

	return ItemsController{
		usecases: usecases,
		logger:   logger,
		config:   config,
	}
}

func (u ItemsController) Get(c *gin.Context) {

	ids := c.Param("ids")

	users, err := u.usecases.Get(c, strings.Split(ids, ","))

	if err != nil {
		u.logger.Error(err.Error(), "get items error")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

func (u ItemsController) Set(c *gin.Context) {

	err := u.usecases.Set(c, u.config.MakvesUrl)

	if err != nil {
		u.logger.Error(err.Error(), "set items error")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, "OK")
}
