package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewClient(amqpURL string) (*Client, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Client{Conn: conn, Ch: ch}, nil
}

func (c *Client) Close() {
	if c.Ch != nil {
		c.Ch.Close()
	}
	if c.Conn != nil {
		c.Conn.Close()
	}
}

func (c *Client) DeclareQueue(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return c.Ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (c *Client) Publish(ctx context.Context, exchange, routingKey string, mandatory, immediate bool, msg amqp.Publishing) error {
	return c.Ch.PublishWithContext(ctx, exchange, routingKey, mandatory, immediate, msg)
}

func (c *Client) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return c.Ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}
