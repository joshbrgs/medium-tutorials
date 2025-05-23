package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/joshbrgs/flipping-out/internal/services"
	"github.com/segmentio/kafka-go"
)

type FeatureEvent struct {
	Value bool `json:"value"`
}

func StartConsumer(h *services.Hub) {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka.default.svc.cluster.local:9092"},
		Topic:    "go-feature-flag-events",
		GroupID:  "feature-analytics",
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  1 * time.Second,
	})
	defer reader.Close()

	log.Println("Starting Kafka consumer...")

	for {
		select {
		case <-h.Done:
			log.Println("Stopping Kafka consumer...")
			return
		default:
			ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)

			m, err := reader.ReadMessage(ctxWithTimeout)
			cancel()

			if err != nil {
				log.Printf("Kafka read error: %v", err)
				time.Sleep(time.Second)
				continue
			}

			var event FeatureEvent
			if err := json.Unmarshal(m.Value, &event); err != nil {
				log.Printf("JSON unmarshal error: %v", err)
				continue
			}

			// Rate limit broadcasts to prevent spam
			h.BroadcastMu.Lock()
			if time.Since(h.LastBroadcast) > 100*time.Millisecond {
				h.LastBroadcast = time.Now()
				h.BroadcastMu.Unlock()

				log.Printf("Broadcasting event: %+v", event)
				h.Broadcast(event)
			} else {
				h.BroadcastMu.Unlock()
				// Skip this event due to rate limiting
			}
		}
	}
}
