package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/yandex-taxi/server/genprotos"
	"github.com/ruziba3vich/yandex-taxi/server/internal/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	CoordinatesService struct {
		storage *storage.TaxiCoordinatesStorage
		logger  *log.Logger
		*genprotos.UnimplementedTaxiCoordinatesServiceServer
	}
)

func NewCoordinatesStorage(storage *storage.TaxiCoordinatesStorage, logger *log.Logger) *CoordinatesService {
	return &CoordinatesService{
		storage: storage,
		logger:  logger,
	}
}

func (c *CoordinatesService) GetTaxiCoordinatesById(ctx context.Context, req *genprotos.GetTaxiCoordinatesByIdRequest) (*emptypb.Empty, error) {
	c.logger.Println("-- RECEIVED A REQUEST INTO GetTaxiCoordinatesById SERVICE --")
	return c.storage.GetTaxiCoordinatesById(ctx, req)
}
