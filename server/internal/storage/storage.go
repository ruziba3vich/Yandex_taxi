package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/ruziba3vich/yandex-taxi/server/genprotos"
	"golang.org/x/sync/errgroup"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type (
	TaxiCoordinatesStorage struct {
		producer sarama.AsyncProducer
		logger   *log.Logger
		rander   *rand.Rand
		db       *DB
		ticker   *time.Ticker
		wg       *sync.WaitGroup
		errgr    *errgroup.Group
		stop     chan byte
	}
)

func NewTaxiCoordinatesStorage(producer sarama.AsyncProducer, logger *log.Logger, rander *rand.Rand, db *DB, ticker *time.Ticker, wg *sync.WaitGroup, errgr *errgroup.Group, stop chan byte) *TaxiCoordinatesStorage {
	return &TaxiCoordinatesStorage{
		producer: producer,
		logger:   logger,
		rander:   rander,
		db:       db,
		ticker:   ticker,
		wg:       wg,
		errgr:    errgr,
		stop:     stop,
	}
}

func (t *TaxiCoordinatesStorage) GetTaxiCoordinatesById(ctx context.Context, req *genprotos.GetTaxiCoordinatesByIdRequest) (*emptypb.Empty, error) {
	if ok := t.db.Get(req.TaxiId); !ok {
		return nil, fmt.Errorf("taxi with %s id is not having a drive yer or not exists", req.TaxiId)
	}
	for {
		select {
		case <-t.ticker.C:
			t.wg.Add(1)
			t.errgr.Go(func() error {
				for {
					select {
					case success := <-t.producer.Successes():
						t.logger.Printf("Message is stored in partition %d, offset %d\n", success.Partition, success.Offset)
						return nil
					case err := <-t.producer.Errors():
						t.logger.Printf("Failed to produce message: %v\n", err.Err)
						return err
					}
				}
			})

			data := t.generateRandomLocation(req.TaxiId)
			byteData, err := json.Marshal(data)
			if err != nil {
				t.logger.Printf("-- ERROR OCCURED WHILE MARSHALING DATA %v --\n", data)
				return nil, err
			}
			msg := &sarama.ProducerMessage{
				Topic: "taxi-location-updates",
				Key:   sarama.StringEncoder(req.TaxiId),
				Value: sarama.ByteEncoder(byteData),
			}
			t.producer.Input() <- msg
			if err := t.errgr.Wait(); err != nil {
				return nil, err
			}
			t.wg.Done()
		case <-t.stop:
			t.wg.Wait()
			return nil, nil
		}
	}
}

func (t *TaxiCoordinatesStorage) generateRandomLocation(taxiID string) *genprotos.TaxiCoordinate {
	return &genprotos.TaxiCoordinate{
		TaxiId: taxiID,
		Angle:  float32(t.rander.Int32() % 180),
		X:      float32(t.rander.Int32() % 1_000),
		Y:      float32(t.rander.Int32() % 1_000),
	}
}
