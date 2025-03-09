package websocket

type (
	AggTradeHandler                func(symbol string, event *WsAggTradeEvent)
	TradeHandler                   func(symbol string, event *WsTradeEvent)
	KlineHandler                   func(symbol string, event *WsKlineEvent)
	MiniTickerHandler              func(symbol string, event *WsMiniTickerEvent)
	TickerHandler                  func(symbol string, event *WsTickerEvent)
	RollingWindowTickerHandler     func(symbol string, event *WsRollingWindowTickerEvent)
	BookTickerHandler              func(symbol string, event *WsBookTickerEvent)
	AvgPriceHandler                func(symbol string, event *WsAvgPriceEvent)
	DepthHandler                   func(symbol string, event *WsDepthEvent)
	DepthUpdateHandler             func(symbol string, event *WsDepthUpdateEvent)
	AllMiniTickersHandler          func(events []*WsMiniTickerEvent)
	AllTickersHandler              func(events []*WsTickerEvent)
	AllRollingWindowTickersHandler func(events []*WsRollingWindowTickerEvent)
)
