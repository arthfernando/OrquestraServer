package routes

import (
	"painellembretes/models"
	"painellembretes/rabbitmq"
	"painellembretes/reminder"
	"painellembretes/sse"

	"github.com/gin-gonic/gin"
)

var (
	ch = make(chan models.Reminder)
)

func SetRoutes() {
	go rabbitmq.ConsumeMessage(ch)
	r := gin.Default()

	r.POST("/send", reminder.SendReminder)
	r.GET("/event-stream", sse.GetEventStream(ch))

	r.Run(":3000")
}
