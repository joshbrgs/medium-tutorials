package services

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type FeatureEvent struct {
	Value bool `json:"value"`
}

type Hub struct {
	clients       map[*websocket.Conn]bool
	register      chan *websocket.Conn
	unregister    chan *websocket.Conn
	broadcast     chan interface{}
	mu            sync.Mutex
	Done          chan struct{}
	LastBroadcast time.Time
	BroadcastMu   sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		broadcast:  make(chan interface{}),
		Done:       make(chan struct{}),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.mu.Lock()
			h.clients[conn] = true
			h.mu.Unlock()
			log.Printf("Client connected. Total clients: %d", len(h.clients))

		case conn := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
			h.mu.Unlock()
			log.Printf("Client disconnected. Total clients: %d", len(h.clients))

		case message := <-h.broadcast:
			h.mu.Lock()
			clientCount := len(h.clients)
			if clientCount == 0 {
				h.mu.Unlock()
				continue
			}

			for conn := range h.clients {
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				if err := conn.WriteJSON(message); err != nil {
					log.Printf("WebSocket write error: %v", err)
					conn.Close()
					delete(h.clients, conn)
				}
			}
			h.mu.Unlock()

		case <-h.Done:
			return
		}
	}
}

func (h *Hub) Broadcast(message interface{}) {
	select {
	case h.broadcast <- message:
	case <-time.After(5 * time.Second):
		log.Println("Broadcast timeout - hub may be blocked")
	}
}

func (h *Hub) Register(conn *websocket.Conn) {
	h.register <- conn
}

func (h *Hub) Unregister(conn *websocket.Conn) {
	h.unregister <- conn
}

func (h *Hub) Stop() {
	close(h.Done)
}
