package main

import (
	"fmt"
	"random-quotes/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {

	repository := app.NewRepository()
	service := app.NewService(repository)
	handler := app.NewHandler(service)

	router := gin.Default()
	router.GET("/", handler.Quotes.HandleRandomQuotes)

	port := 8080
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server: %v", err))
	}
}
