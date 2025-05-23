package v1

import (
	"log"
	"net/http"

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
	return &WebsocketController{hub: hub}
}

func (wc *WebsocketController) HandleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wc.hub.Register(conn)

	go func() {
		defer wc.hub.Unregister(conn)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()

	go kafka.StartConsumer(c, wc.hub)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
