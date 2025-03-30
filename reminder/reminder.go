package reminder

import (
	"net/http"
	"painellembretes/models"
	"painellembretes/rabbitmq"

	"github.com/gin-gonic/gin"
)

func SendReminder(c *gin.Context) {
	var reminder models.Reminder

	err := c.ShouldBind(&reminder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "[ERR] Binding body request",
			"error":   err.Error(),
		})
	}

	err = rabbitmq.SendMessage(reminder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "[ERR] Sending reminder",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reminder sent",
	})
}

func ReceiveReminder(c *gin.Context) {
	rabbitmq.ConsumeMessage()
}
