package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ruziba3vich/yandex-taxi/server/genprotos"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	TaxiCreationStorage struct {
		db *DB
	}
)

func NewTaxiCreatorStorage(db *DB) *TaxiCreationStorage {
	return &TaxiCreationStorage{
		db: db,
	}
}

func (t *TaxiCreationStorage) CreateTaxi(context.Context, *emptypb.Empty) (*genprotos.CreateTaxiResponse, error) {
	id := uuid.New().String()
	if ok := t.db.Add(id); !ok {
		return nil, fmt.Errorf("something went wrong")
	}
	return &genprotos.CreateTaxiResponse{
		TaxiId: id,
	}, nil
}
