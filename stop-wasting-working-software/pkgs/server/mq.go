package server

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Config struct holds RabbitMQ connection configuration
type Config struct {
	URL string
}

// ConnectionOption defines the options pattern for setting connection configuration
type ConnectionOption func(*Config) error

// WithRabbitMQURL sets the RabbitMQ URL in the config
func WithRabbitMQURL(url string) ConnectionOption {
	return func(c *Config) error {
		c.URL = url
		return nil
	}
}

// RabbitMQ holds the connection and channel
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ initializes a RabbitMQ connection and channel with the provided options
func NewRabbitMQ(opts ...ConnectionOption) (*RabbitMQ, error) {
	config := &Config{
		URL: "amqp://guest:guest@localhost:5672/", // default URL
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	// Establish connection
	conn, err := amqp.Dial(config.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Create a channel
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

// Close gracefully closes the RabbitMQ connection and channel
func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

// Publish sends a message to the specified queue
func (r *RabbitMQ) Publish(ctx context.Context, queueName, body string) error {
	_, err := r.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = r.channel.PublishWithContext(
		ctx,
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Message published to queue %s: %s", queueName, body)
	return nil
}

// Consume consumes messages from the specified queue
func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	_, err := r.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	messages, err := r.channel.Consume(
		queueName, // queue name
		"",        // consumer tag
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}

	return messages, nil
}
