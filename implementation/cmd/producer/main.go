package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"keda-worker/internal/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	count := flag.Int("count", 10, "number of messages to send")
	flag.Parse()

	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		amqpURL = "amqp://user:PASSWORD@localhost:5672/"
	}

	client, err := rabbitmq.NewClient(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer client.Close()

	q, err := client.DeclareQueue(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	for i := range *count {
		body := fmt.Sprintf("Message %d", i)
		err := client.Publish(
			context.Background(),
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			})
		if err != nil {
			log.Printf("Failed to publish a message: %s", err)
		}
		log.Printf(" [x] Sent %s\n", body)
	}
}
