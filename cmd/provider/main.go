package main

import (
	"net/http"

	rabbitMQConfig "github.com/achmadAlli/sample-go-echo-rabbitmq/config"
	"github.com/labstack/echo"
	"github.com/streadway/amqp"
)

type service struct {
	rabbitmq *rabbitMQConfig.RabbitMQClient
}

func main() {
	e := echo.New()

	s := service{
		rabbitmq: rabbitMQConfig.Connect(),
	}

	// initial route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "ok",
		})
	})

	e.GET("/send_1", s.sendToFirstConsumer)
	e.GET("/send_2", s.sendToSecondConsumer)
	e.GET("/send_3", s.sendToBothConsumer)

	e.Start(":8000")
}

func (s service) sendToFirstConsumer(c echo.Context) error {
	s.rabbitmq.DeclareQueue("FIRST_QUEUE")
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("message for first consumer"),
	}

	err := s.rabbitmq.PublishMessage("FIRST_QUEUE", message)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "message send to first consumer",
	})
}

func (s service) sendToSecondConsumer(c echo.Context) error {
	s.rabbitmq.DeclareQueue("SECOND_QUEUE")

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("message for second consumer"),
	}

	s.rabbitmq.PublishMessage("SECOND_QUEUE", message)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "message send to second consumer",
	})
}

func (s service) sendToBothConsumer(c echo.Context) error {
	s.rabbitmq.DeclareQueue("BOTH_QUEUE")

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("message for both consumer"),
	}

	s.rabbitmq.PublishMessage("BOTH_QUEUE", message)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "message send to both consumer",
	})
}
