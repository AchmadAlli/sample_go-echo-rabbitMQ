package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	channel *amqp.Channel
}

func Connect() *RabbitMQClient {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Missing env data")
	}

	url := os.Getenv("AMQP_SERVER_URL")

	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return &RabbitMQClient{channel: channel}
}

func (r *RabbitMQClient) PublishMessage(queueName string, message amqp.Publishing) error {
	return r.channel.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		message,   // message to publish
	)
}

func (r *RabbitMQClient) Consume(queueName string) (<-chan amqp.Delivery, error) {
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

func (r RabbitMQClient) DeclareQueue(queueName string) {
	_, err := r.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}
}
