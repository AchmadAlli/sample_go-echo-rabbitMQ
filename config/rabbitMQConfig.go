package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	channel *amqp.Channel
}

func Connect() *rabbitMQ {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Missing env data")
	}

	url := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	return &rabbitMQ{channel: channelRabbitMQ}
}

func (r rabbitMQ) PublishMessage(queueName string, message amqp.Publishing) {
	r.channel.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		message,   // message to publish
	)
}

func (r rabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	messages, err := r.channel.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // arguments
	)

	return messages, err
}
