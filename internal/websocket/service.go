package websocket

import (
	"context"
	"sync"

	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/agniBit/cryptonian/model/websocket"
	ws "github.com/fasthttp/websocket"
)

type BinanceWebSocket struct {
	ctx              context.Context
	cfg              *cfg.Config
	conn             *ws.Conn
	url              string
	subscribed       map[string]bool
	mu               sync.Mutex
	send             chan []byte
	reconnectAttempt int

	aggTradeHandler                websocket.AggTradeHandler
	tradeHandler                   websocket.TradeHandler
	klineHandler                   websocket.KlineHandler
	miniTickerHandler              websocket.MiniTickerHandler
	allMiniTickersHandler          websocket.AllMiniTickersHandler
	tickerHandler                  websocket.TickerHandler
	allTickersHandler              websocket.AllTickersHandler
	rollingWindowTickerHandler     websocket.RollingWindowTickerHandler
	allRollingWindowTickersHandler websocket.AllRollingWindowTickersHandler
	bookTickerHandler              websocket.BookTickerHandler
	avgPriceHandler                websocket.AvgPriceHandler
	depthHandler                   websocket.DepthHandler
	depthUpdateHandler             websocket.DepthUpdateHandler
}

func NewBinanceWebSocket(ctx context.Context, cfg *cfg.Config) *BinanceWebSocket {
	return &BinanceWebSocket{
		ctx:        ctx,
		cfg:        cfg,
		url:        cfg.Binance.WsURL,
		subscribed: make(map[string]bool),
		send:       make(chan []byte, 100),
	}
}

// Handler registration methods
func (ws *BinanceWebSocket) AddAggTradeHandler(handler websocket.AggTradeHandler) {
	ws.aggTradeHandler = handler
}

func (ws *BinanceWebSocket) AddTradeHandler(handler websocket.TradeHandler) {
	ws.tradeHandler = handler
}

func (ws *BinanceWebSocket) AddKlineHandler(handler websocket.KlineHandler) {
	ws.klineHandler = handler
}

func (ws *BinanceWebSocket) AddMiniTickerHandler(handler websocket.MiniTickerHandler) {
	ws.miniTickerHandler = handler
}

func (ws *BinanceWebSocket) AddAllMiniTickersHandler(handler websocket.AllMiniTickersHandler) {
	ws.allMiniTickersHandler = handler
}

func (ws *BinanceWebSocket) AddTickerHandler(handler websocket.TickerHandler) {
	ws.tickerHandler = handler
}

func (ws *BinanceWebSocket) AddAllTickersHandler(handler websocket.AllTickersHandler) {
	ws.allTickersHandler = handler
}

func (ws *BinanceWebSocket) AddRollingWindowTickerHandler(handler websocket.RollingWindowTickerHandler) {
	ws.rollingWindowTickerHandler = handler
}

func (ws *BinanceWebSocket) AddAllRollingWindowTickersHandler(handler websocket.AllRollingWindowTickersHandler) {
	ws.allRollingWindowTickersHandler = handler
}

func (ws *BinanceWebSocket) AddBookTickerHandler(handler websocket.BookTickerHandler) {
	ws.bookTickerHandler = handler
}

func (ws *BinanceWebSocket) AddAvgPriceHandler(handler websocket.AvgPriceHandler) {
	ws.avgPriceHandler = handler
}

func (ws *BinanceWebSocket) AddDepthHandler(handler websocket.DepthHandler) {
	ws.depthHandler = handler
}

func (ws *BinanceWebSocket) AddDepthUpdateHandler(handler websocket.DepthUpdateHandler) {
	ws.depthUpdateHandler = handler
}
