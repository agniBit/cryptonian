package websocket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/fasthttp/websocket"
)

type Request struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

type CombinedStreamMessage struct {
	Stream string          `json:"stream"`
	Data   json.RawMessage `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type Response struct {
	Result interface{}    `json:"result"`
	Error  *ErrorResponse `json:"error"`
	ID     int            `json:"id"`
}

func (ws *BinanceWebSocket) Start() error {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(ws.url, http.Header{})
	if err != nil {
		return err
	}

	ws.conn = conn
	go ws.readLoop()
	go ws.writeLoop()
	go ws.keepAlive()

	if len(ws.subscribed) > 0 {
		ws.resubscribe()
	}

	return nil
}

func (ws *BinanceWebSocket) resubscribe() {
	streams := make([]string, 0, len(ws.subscribed))
	for stream := range ws.subscribed {
		streams = append(streams, stream)
	}

	req := Request{
		Method: "SUBSCRIBE",
		Params: streams,
		ID:     generateID(),
	}

	msg, err := json.Marshal(req)
	if err != nil {
		logger.Error(context.Background(), "Error marshaling resubscribe request", err)
		return
	}

	ws.send <- msg
}

func (ws *BinanceWebSocket) readLoop() {
	defer ws.reconnect()

	for {
		select {
		case <-ws.ctx.Done():
			return
		default:
			_, message, err := ws.conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}

			ws.handleMessage(message)
		}
	}
}

func (ws *BinanceWebSocket) writeLoop() {
	rateLimiter := time.NewTicker(200 * time.Millisecond)
	defer rateLimiter.Stop()

	for {
		select {
		case <-ws.ctx.Done():
			return
		case msg := <-ws.send:
			<-rateLimiter.C
			ws.mu.Lock()
			err := ws.conn.WriteMessage(websocket.TextMessage, msg)
			ws.mu.Unlock()
			if err != nil {
				logger.Error(context.Background(), "Write error", err)
				return
			}
		}
	}
}

func (ws *BinanceWebSocket) keepAlive() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ws.ctx.Done():
			return
		case <-ticker.C:
			ws.mu.Lock()
			err := ws.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second))
			ws.mu.Unlock()
			if err != nil {
				logger.Error(context.Background(), "Keep alive error", err)
				return
			}
		}
	}
}

func (ws *BinanceWebSocket) Subscribe(streams []string) {
	req := Request{
		Method: "SUBSCRIBE",
		Params: streams,
		ID:     generateID(),
	}

	msg, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error marshaling subscribe request: %v", err)
		return
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()
	for _, stream := range streams {
		ws.subscribed[stream] = true
	}

	ws.send <- msg
}

func (ws *BinanceWebSocket) Unsubscribe(streams []string) {
	req := Request{
		Method: "UNSUBSCRIBE",
		Params: streams,
		ID:     generateID(),
	}

	msg, err := json.Marshal(req)
	if err != nil {
		logger.Error(context.Background(), "Error marshaling unsubscribe request", err)
		return
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()
	for _, stream := range streams {
		delete(ws.subscribed, stream)
	}

	ws.send <- msg
}

func (ws *BinanceWebSocket) reconnect() {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	select {
	case <-ws.ctx.Done():
		return
	default:
	}

	logger.Warn(context.Background(), "Reconnecting to Binance WebSocket", nil)
	ws.conn.Close()
	ws.reconnectAttempt++

	backoff := time.Duration(ws.reconnectAttempt*2) * time.Second
	if backoff > 60*time.Second {
		backoff = 60 * time.Second
	}

	time.Sleep(backoff)

	if err := ws.Start(); err != nil {
		logger.Error(context.Background(), "Error reconnecting to Binance WebSocket", err)
		go ws.reconnect()
	} else {
		ws.reconnectAttempt = 0
	}
}

func (ws *BinanceWebSocket) Close() {
	ws.conn.Close()
}

var idCounter int

func generateID() int {
	idCounter++
	return idCounter
}
