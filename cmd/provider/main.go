package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// initial route
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "ok",
		})
	})

	e.GET("/send_1", sendToFirstConsumer)
	e.GET("/send_2", sendToSecondConsumer)
	e.GET("/send_3", sendToBothConsumer)

	e.Start(":8000")
}

func sendToFirstConsumer(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"message": "message send to first consumer",
	})
}

func sendToSecondConsumer(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "message send to second consumer",
	})
}

func sendToBothConsumer(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "message send to both consumer",
	})
}
