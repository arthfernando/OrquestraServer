package routes

import (
	"painellembretes/reminder"

	"github.com/gin-gonic/gin"
)

func SetRoutes() {
	r := gin.Default()

	r.POST("/send", reminder.SendReminder)
	r.GET("/", reminder.ReceiveReminder)

	r.Run(":3000")
}
