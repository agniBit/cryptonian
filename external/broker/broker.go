package broker

import (
	"context"
	"errors"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/model/cfg"
)

type brokerType int

const (
	Binance brokerType = iota
)

type Broker interface {
}

type broker struct {
	brokerSyncMap map[brokerType]BrokerService
}

func NewBrokerService(cfg *cfg.Config) Broker {
	// init all the broker services

	brokerSyncMap := map[brokerType]BrokerService{}

	return &broker{
		brokerSyncMap: brokerSyncMap,
	}
}

type BrokerService interface {
}

func (b *broker) getBrokerService(ctx context.Context, brokerType brokerType) (Broker, error) {
	brokerService, ok := b.brokerSyncMap[brokerType]
	if !ok {
		err := errors.New("Broker not found")
		logger.Error(ctx, "Broker not found", err)
		return nil, err
	}

	return brokerService, nil
}
