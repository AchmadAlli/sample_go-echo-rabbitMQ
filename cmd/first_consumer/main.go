package main

import (
	"log"

	rabbitMQConfig "github.com/achmadAlli/sample-go-echo-rabbitmq/config"
)

type service struct {
	rabbitmq *rabbitMQConfig.RabbitMQClient
}

func main() {

	s := service{rabbitmq: rabbitMQConfig.Connect()}

	messages, err := s.rabbitmq.Consume("FIRST_QUEUE")
	if err != nil {
		log.Println(err)
	}

	// make channel to create infinite loop
	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("# message from FIRST_QUEUES : %s\n", message.Body)
		}
	}()

	<-forever

}
