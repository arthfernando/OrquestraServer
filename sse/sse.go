package sse

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"painellembretes/models"

	"github.com/gin-gonic/gin"
)

func PostEventStream(ch chan string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.EventStreamRequest
		if err := c.ShouldBind(&request); err != nil {
			errorMessage := fmt.Sprintf("request validation error: %s", err.Error())
			c.JSON(http.StatusBadRequest, models.SendFailResponse[string]{
				Status: "fail",
				Data:   errors.New(errorMessage).Error(),
			})

			return
		}

		ch <- request.Message
		c.JSON(http.StatusCreated, models.SendSuccessResponse[interface{}]{
			Status: "success",
			Data:   &request.Message,
		})
	}

}

func GetEventStream(ch chan string) gin.HandlerFunc {
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
