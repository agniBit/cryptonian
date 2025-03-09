package websocket

import (
	"encoding/json"
	"strings"

	"github.com/agniBit/cryptonian/model/websocket"
)

func (ws *BinanceWebSocket) handleMessage(msg []byte) {
	var combinedEvent struct {
		Stream string          `json:"stream"`
		Data   json.RawMessage `json:"data"`
	}

	err := json.Unmarshal(msg, &combinedEvent)
	if err == nil {
		// Handle combined stream format
		streamParts := strings.Split(combinedEvent.Stream, "@")
		symbol := strings.ToUpper(streamParts[0])

		var dataEvent struct{ E string }
		json.Unmarshal(combinedEvent.Data, &dataEvent)

		switch dataEvent.E {
		case "aggTrade":
			var event websocket.WsAggTradeEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.aggTradeHandler != nil {
				ws.aggTradeHandler(symbol, &event)
			}

		case "trade":
			var event websocket.WsTradeEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.tradeHandler != nil {
				ws.tradeHandler(symbol, &event)
			}

		case "kline":
			var event websocket.WsKlineEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.klineHandler != nil {
				ws.klineHandler(symbol, &event)
			}

		case "24hrMiniTicker":
			var event websocket.WsMiniTickerEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.miniTickerHandler != nil {
				ws.miniTickerHandler(symbol, &event)
			}

		case "!miniTicker@arr":
			var events []*websocket.WsMiniTickerEvent
			json.Unmarshal(combinedEvent.Data, &events)
			if ws.allMiniTickersHandler != nil {
				ws.allMiniTickersHandler(events)
			}

		case "24hrTicker":
			var event websocket.WsTickerEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.tickerHandler != nil {
				ws.tickerHandler(symbol, &event)
			}

		case "!ticker@arr":
			var events []*websocket.WsTickerEvent
			json.Unmarshal(combinedEvent.Data, &events)
			if ws.allTickersHandler != nil {
				ws.allTickersHandler(events)
			}

		case "1hTicker", "4hTicker", "1dTicker":
			var event websocket.WsRollingWindowTickerEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.rollingWindowTickerHandler != nil {
				ws.rollingWindowTickerHandler(symbol, &event)
			}

		case "!ticker_1h@arr", "!ticker_4h@arr", "!ticker_1d@arr":
			var events []*websocket.WsRollingWindowTickerEvent
			json.Unmarshal(combinedEvent.Data, &events)
			if ws.allRollingWindowTickersHandler != nil {
				ws.allRollingWindowTickersHandler(events)
			}

		case "bookTicker":
			var event websocket.WsBookTickerEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.bookTickerHandler != nil {
				ws.bookTickerHandler(symbol, &event)
			}

		case "avgPrice":
			var event websocket.WsAvgPriceEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.avgPriceHandler != nil {
				ws.avgPriceHandler(symbol, &event)
			}

		case "depthUpdate":
			var event websocket.WsDepthUpdateEvent
			json.Unmarshal(combinedEvent.Data, &event)
			if ws.depthUpdateHandler != nil {
				ws.depthUpdateHandler(symbol, &event)
			}

		default:
			// Handle partial book depth stream
			var depthEvent websocket.WsDepthEvent
			json.Unmarshal(combinedEvent.Data, &depthEvent)
			if depthEvent.LastUpdateID != 0 && ws.depthHandler != nil {
				ws.depthHandler(symbol, &depthEvent)
			}
		}
	} else {
		// Handle raw stream format
		var dataEvent struct{ E string }
		json.Unmarshal(msg, &dataEvent)

		switch dataEvent.E {
		case "aggTrade":
			var event websocket.WsAggTradeEvent
			json.Unmarshal(msg, &event)
			if ws.aggTradeHandler != nil {
				ws.aggTradeHandler(event.Symbol, &event)
			}

		case "trade":
			var event websocket.WsTradeEvent
			json.Unmarshal(msg, &event)
			if ws.tradeHandler != nil {
				ws.tradeHandler(event.Symbol, &event)
			}

		case "kline":
			var event websocket.WsKlineEvent
			json.Unmarshal(msg, &event)
			if ws.klineHandler != nil {
				ws.klineHandler(event.Symbol, &event)
			}

		case "24hrMiniTicker":
			var event websocket.WsMiniTickerEvent
			json.Unmarshal(msg, &event)
			if ws.miniTickerHandler != nil {
				ws.miniTickerHandler(event.Symbol, &event)
			}

		case "!miniTicker@arr":
			var events []*websocket.WsMiniTickerEvent
			json.Unmarshal(msg, &events)
			if ws.allMiniTickersHandler != nil {
				ws.allMiniTickersHandler(events)
			}

		case "24hrTicker":
			var event websocket.WsTickerEvent
			json.Unmarshal(msg, &event)
			if ws.tickerHandler != nil {
				ws.tickerHandler(event.Symbol, &event)
			}

		case "!ticker@arr":
			var events []*websocket.WsTickerEvent
			json.Unmarshal(msg, &events)
			if ws.allTickersHandler != nil {
				ws.allTickersHandler(events)
			}

		case "1hTicker", "4hTicker", "1dTicker":
			var event websocket.WsRollingWindowTickerEvent
			json.Unmarshal(msg, &event)
			if ws.rollingWindowTickerHandler != nil {
				ws.rollingWindowTickerHandler(event.Symbol, &event)
			}

		case "!ticker_1h@arr", "!ticker_4h@arr", "!ticker_1d@arr":
			var events []*websocket.WsRollingWindowTickerEvent
			json.Unmarshal(msg, &events)
			if ws.allRollingWindowTickersHandler != nil {
				ws.allRollingWindowTickersHandler(events)
			}

		case "bookTicker":
			var event websocket.WsBookTickerEvent
			json.Unmarshal(msg, &event)
			if ws.bookTickerHandler != nil {
				ws.bookTickerHandler(event.Symbol, &event)
			}

		case "avgPrice":
			var event websocket.WsAvgPriceEvent
			json.Unmarshal(msg, &event)
			if ws.avgPriceHandler != nil {
				ws.avgPriceHandler(event.Symbol, &event)
			}

		case "depthUpdate":
			var event websocket.WsDepthUpdateEvent
			json.Unmarshal(msg, &event)
			if ws.depthUpdateHandler != nil {
				ws.depthUpdateHandler(event.Symbol, &event)
			}

		default:
			// Handle raw partial book depth (symbol unknown)
			var depthEvent websocket.WsDepthEvent
			json.Unmarshal(msg, &depthEvent)
			if depthEvent.LastUpdateID != 0 && ws.depthHandler != nil {
				// Symbol unknown for raw depth streams - handle with care
				ws.depthHandler("", &depthEvent)
			}
		}
	}
}
