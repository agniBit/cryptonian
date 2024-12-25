package broker

import (
	"context"
	"errors"
	"sync"

	"github.com/agniBit/cryptonian/external/broker/binance"
	"github.com/agniBit/cryptonian/model/cfg"
)

const (
	Binance = 1
)

type Broker interface {
}

type broker struct {
	brokerSyncMap *sync.Map
}

func NewBrokerService(cfg *cfg.Config) Broker {
	// init all the broker services
	binace := binance.NewBinanceService(cfg)

	brokerSyncMap := &sync.Map{}
	brokerSyncMap.Store(Binance, binace)

	return &broker{
		brokerSyncMap: brokerSyncMap,
	}
}

func (b *broker) getBrokerService(ctx context.Context, brokerId int64) (Broker, error) {
	if broker, ok := b.brokerSyncMap.Load(brokerId); ok {
		return broker.(Broker), nil
	}
	return nil, errors.New("broker not found")
}
