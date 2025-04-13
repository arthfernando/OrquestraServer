package models

type EventStreamRequest struct {
	Message string `form:"message" json:"message" binding:"required,max=100"`
}

type SendFailResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data,omitempty"`
}

type SendSuccessResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data,omitempty"`
}
