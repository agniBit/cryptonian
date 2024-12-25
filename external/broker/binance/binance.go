package binance

import (
	"context"
	"sync"

	"github.com/agniBit/cryptonian/model/cfg"
)

type binance struct {
	cfg           *cfg.Config
	userClientMap *sync.Map
}

func NewBinanceService(cfg *cfg.Config) *binance {
	return &binance{
		cfg:           cfg,
		userClientMap: &sync.Map{},
	}
}

func (b *binance) RegisterNewWsAggTradeServe(ctx context.Context) {}
