package main

import (
	"log"
	"os"
	"time"

	"keda-worker/internal/rabbitmq"
)

func main() {
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

	err = client.Ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}

	msgs, err := client.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack is false for manual acknowledgement
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	log.Println("Worker started. Waiting for messages.")

	// Process messages in a loop
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// Simulate work
			time.Sleep(1 * time.Second)
			log.Printf("Done processing message")
			d.Ack(false) // Acknowledge the message
		}
	}()
	<-forever // Block forever
}
