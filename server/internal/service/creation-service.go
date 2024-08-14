package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/yandex-taxi/server/genprotos"
	"github.com/ruziba3vich/yandex-taxi/server/internal/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	CreationServce struct {
		storage *storage.TaxiCreationStorage
		logger  *log.Logger
		*genprotos.UnimplementedTaxiCreationServiceServer
	}
)

func NewCreationService(storage *storage.TaxiCreationStorage, logger *log.Logger) *CreationServce {
	return &CreationServce{
		storage: storage,
		logger:  logger,
	}
}

func (c *CreationServce) CreateTaxi(ctx context.Context, e *emptypb.Empty) (*genprotos.CreateTaxiResponse, error) {
	c.logger.Println("-- GOT A REQUEST INTO CreateTaxi SERVICE --")
	return c.storage.CreateTaxi(ctx, e)
}
