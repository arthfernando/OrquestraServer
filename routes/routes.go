package routes

import (
	"painellembretes/reminder"
	"painellembretes/sse"

	"github.com/gin-gonic/gin"
)

var (
	ch = make(chan string)
)

func SetRoutes() {
	r := gin.Default()

	r.POST("/send", reminder.SendReminder)
	// r.GET("/", reminder.ReceiveReminder)
	r.POST("/event-stream", sse.PostEventStream(ch))
	r.GET("/event-stream", sse.GetEventStream(ch))

	r.Run(":3000")
}
