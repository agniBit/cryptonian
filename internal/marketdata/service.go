package marketdata

import (
	"context"
	"time"

	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/patrickmn/go-cache"
)

type service struct {
	klineDataCache *cache.Cache
	aggTradeCache  *cache.Cache
	depthCache     *cache.Cache
}

type ServiceInterface interface {
	UpdateKlineData(ctx context.Context, symbol string, interval string) error
	UpdateAggTradeData(ctx context.Context, symbol string, interval string) error
	UpdateDepthData(ctx context.Context, symbol string) error
}

func NewMarketDataService(cfg *cfg.Config) *service {
	return &service{
		klineDataCache: cache.New(-1, -1),
		aggTradeCache:  cache.New(-1, -1),
		depthCache:     cache.New(30*time.Second, 1*time.Minute),
	}
}

func (s *service) UpdateKlineData(ctx context.Context, symbol string, interval string) error {
	return nil
}

func (s *service) UpdateAggTradeData(ctx context.Context, symbol string, interval string) error {
	return nil
}

func (s *service) UpdateDepthData(ctx context.Context, symbol string) error {
	return nil
}
