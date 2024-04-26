package server

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQOption represents a function that modifies RabbitMQ connection options
type RabbitMQOption func(*amqp.Connection)

// WithRabbitMQURL sets the RabbitMQ URL option
func WithRabbitMQURL(url string) RabbitMQOption {
	return func(conn *amqp.Connection) {
		// This function can handle any additional connection configuration
		// For example, setting TLS configuration, credentials, etc.
		// You can parse the URL and configure the connection accordingly
	}
}

// WithQueue declares a queue with the specified name and options
func WithQueue(name string, options amqp.Table) RabbitMQOption {
	return func(conn *amqp.Connection) {
		ch, err := conn.Channel()
		if err != nil {
			// Handle error
			return
		}
		// Declare the queue on the connection
		// You can pass additional options such as durable, exclusive, autoDelete, etc.
		_, err = ch.QueueDeclare(name, false, false, false, false, options)
		if err != nil {
			// Handle error
		}
	}
}

// WithExchange declares an exchange with the specified name, type, and options
func WithExchange(name, kind string, options amqp.Table) RabbitMQOption {
	return func(conn *amqp.Connection) {
		ch, err := conn.Channel()
		if err != nil {
			// Handle error
			return
		}
		// Declare the exchange on the connection
		// You can pass additional options such as durable, autoDelete, etc.
		if err = ch.ExchangeDeclare(name, kind, false, false, false, false, options); err != nil {
			// Handle error
		}
	}
}

// WithPublish confirms publishing to a specific exchange with the provided routing key
func WithPublish(exchange, routingKey string) RabbitMQOption {
	return func(conn *amqp.Connection) {
		// You can implement logic here to publish to the specified exchange with the provided routing key
		// You may want to return a function that handles the publishing
	}
}

// NewRabbitMQConnection initializes a new connection to RabbitMQ with the provided options
func NewRabbitMQConnection(url string, opts ...RabbitMQOption) (*amqp.Connection, error) {
	// Create a new RabbitMQ connection
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	// Apply provided options
	for _, opt := range opts {
		opt(conn)
	}

	return conn, nil
}
