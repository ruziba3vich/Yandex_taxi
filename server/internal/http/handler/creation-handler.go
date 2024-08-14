package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/yandex-taxi/server/internal/service"
)

type (
	Handler struct {
		service *service.CreationServce
		logger  *log.Logger
	}
)

func NewHandler(service *service.CreationServce, logger *log.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func CreateTaxi(c *gin.Conntext) {
	// service.
}
