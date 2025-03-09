package websocket

type WsAggTradeEvent struct {
	EventTime    int64  `json:"E"`
	Symbol       string `json:"s"`
	AggTradeID   int64  `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeID int64  `json:"f"`
	LastTradeID  int64  `json:"l"`
	TradeTime    int64  `json:"T"`
	IsMaker      bool   `json:"m"`
	M            bool   `json:"M"`
}

type WsTradeEvent struct {
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	TradeID   int64  `json:"t"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
	TradeTime int64  `json:"T"`
	IsMaker   bool   `json:"m"`
	M         bool   `json:"M"`
}

type Kline struct {
	StartTime      int64  `json:"t"`
	CloseTime      int64  `json:"T"`
	Symbol         string `json:"s"`
	Interval       string `json:"i"`
	FirstTradeID   int64  `json:"f"`
	LastTradeID    int64  `json:"L"`
	OpenPrice      string `json:"o"`
	ClosePrice     string `json:"c"`
	HighPrice      string `json:"h"`
	LowPrice       string `json:"l"`
	BaseVolume     string `json:"v"`
	NumberOfTrades int64  `json:"n"`
	IsClosed       bool   `json:"x"`
	QuoteVolume    string `json:"q"`
	TakerBuyBase   string `json:"V"`
	TakerBuyQuote  string `json:"Q"`
	B              string `json:"B"`
}

type WsKlineEvent struct {
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	Kline     Kline  `json:"k"`
}

type WsMiniTickerEvent struct {
	EventTime   int64  `json:"E"`
	Symbol      string `json:"s"`
	ClosePrice  string `json:"c"`
	OpenPrice   string `json:"o"`
	HighPrice   string `json:"h"`
	LowPrice    string `json:"l"`
	BaseVolume  string `json:"v"`
	QuoteVolume string `json:"q"`
}

type WsTickerEvent struct {
	EventTime        int64  `json:"E"`
	Symbol           string `json:"s"`
	PriceChange      string `json:"p"`
	PriceChangePct   string `json:"P"`
	WeightedAvgPrice string `json:"w"`
	FirstTradePrice  string `json:"x"`
	LastPrice        string `json:"c"`
	LastQuantity     string `json:"Q"`
	BestBidPrice     string `json:"b"`
	BestBidQty       string `json:"B"`
	BestAskPrice     string `json:"a"`
	BestAskQty       string `json:"A"`
	OpenPrice        string `json:"o"`
	HighPrice        string `json:"h"`
	LowPrice         string `json:"l"`
	BaseVolume       string `json:"v"`
	QuoteVolume      string `json:"q"`
	OpenTime         int64  `json:"O"`
	CloseTime        int64  `json:"C"`
	FirstTradeID     int64  `json:"F"`
	LastTradeID      int64  `json:"L"`
	TotalTrades      int64  `json:"n"`
}

type WsRollingWindowTickerEvent struct {
	EventTime        int64  `json:"E"`
	Symbol           string `json:"s"`
	PriceChange      string `json:"p"`
	PriceChangePct   string `json:"P"`
	OpenPrice        string `json:"o"`
	HighPrice        string `json:"h"`
	LowPrice         string `json:"l"`
	ClosePrice       string `json:"c"`
	WeightedAvgPrice string `json:"w"`
	BaseVolume       string `json:"v"`
	QuoteVolume      string `json:"q"`
	OpenTime         int64  `json:"O"`
	CloseTime        int64  `json:"C"`
	FirstTradeID     int64  `json:"F"`
	LastTradeID      int64  `json:"L"`
	TotalTrades      int64  `json:"n"`
}

type WsBookTickerEvent struct {
	UpdateID     int64  `json:"u"`
	Symbol       string `json:"s"`
	BestBidPrice string `json:"b"`
	BestBidQty   string `json:"B"`
	BestAskPrice string `json:"a"`
	BestAskQty   string `json:"A"`
}

type WsAvgPriceEvent struct {
	Event         string `json:"e"`
	EventTime     int64  `json:"E"`
	Symbol        string `json:"s"`
	Interval      string `json:"i"`
	AvgPrice      string `json:"w"`
	LastTradeTime int64  `json:"T"`
}

type WsDepthEvent struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type WsDepthUpdateEvent struct {
	EventTime     int64      `json:"E"`
	Symbol        string     `json:"s"`
	FirstUpdateID int64      `json:"U"`
	LastUpdateID  int64      `json:"u"`
	Bids          [][]string `json:"b"`
	Asks          [][]string `json:"a"`
}
