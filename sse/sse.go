package sse

import (
	"io"
	"painellembretes/models"

	"github.com/gin-gonic/gin"
)

func GetEventStream(ch chan models.Reminder) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Stream(func(io.Writer) bool {
			if msg, ok := <-ch; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	}
}
