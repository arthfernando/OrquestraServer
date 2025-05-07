package sse

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func GetEventStream(hub *SSEHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := make(SSEClient)
		hub.add <- client
		defer func() { hub.remove <- client }()

		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		c.Stream(func(io.Writer) bool {
			if msg, ok := <-client; ok {
				log.Println("Emit SSE message")
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	}
}
