package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joshbrgs/flipping-out/internal/kafka"
	"github.com/joshbrgs/flipping-out/internal/services"
)

type WebsocketController struct {
	hub *services.Hub
}

func NewWebsocketController(hub *services.Hub) *WebsocketController {
	go hub.Run()
	go kafka.StartConsumer(hub)
	return &WebsocketController{hub: hub}
}

func (wc *WebsocketController) HandleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	// Set connection limits and timeouts
	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Register connection
	wc.hub.Register(conn)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// Handle websocket messages in goroutine
	go func() {
		defer func() {
			wc.hub.Unregister(conn)
			conn.Close()
		}()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err,
					websocket.CloseGoingAway,
					websocket.CloseAbnormalClosure,
					websocket.CloseNoStatusReceived) {
					log.Printf("WebSocket unexpected close: %v", err)
				}
				break
			}

			switch messageType {
			case websocket.TextMessage:
				log.Printf("Received text message: %s", message)
			case websocket.BinaryMessage:
				log.Printf("Received binary message of length: %d", len(message))
			case websocket.CloseMessage:
				log.Println("Received close message")
				return
			}
		}
	}()

	// Ping loop
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Printf("Ping error: %v", err)
					return
				}
			}
		}
	}()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 45 * time.Second,
}
