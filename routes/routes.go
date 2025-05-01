package routes

import (
	"painellembretes/models"
	"painellembretes/rabbitmq"
	"painellembretes/reminder"
	"painellembretes/sse"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ch = make(chan models.Reminder)
)

func SetRoutes() {
	hub := sse.NewSSEHub()
	r := gin.Default()

	go hub.Run()
	go rabbitmq.ConsumeMessage(hub)

	r.POST("/api/v1/send", reminder.SendReminder)
	r.GET("/api/v1/event-stream", sse.GetEventStream(hub))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "OPTIONS", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Run(":3000")
}
