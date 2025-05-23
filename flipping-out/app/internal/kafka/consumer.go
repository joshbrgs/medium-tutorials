package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/joshbrgs/flipping-out/internal/services"
	"github.com/segmentio/kafka-go"
)

type FeatureEvent struct {
	Value bool `json:"value"`
}

func StartConsumer(ctx context.Context, hub *services.Hub) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka.default.svc.cluster.local:9092"},
		Topic:   "go-feature-flag-events",
		GroupID: "feature-analytics",
	})
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}

		var event FeatureEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("JSON unmarshal error: %v", err)
			continue
		}

		hub.Broadcast(event)
	}
}
