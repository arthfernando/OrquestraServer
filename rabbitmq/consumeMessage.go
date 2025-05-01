package rabbitmq

import (
	"encoding/json"
	"log"
	"painellembretes/config"
	"painellembretes/models"
	"painellembretes/shared"
	"painellembretes/sse"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessage(hub *sse.SSEHub) {
	rabbitmqUrl := config.Get("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitmqUrl)
	shared.FailOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	shared.FailOnError(err, "Failed to open RabbitMQ channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"painellembretes-server",
		false,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "Failed to declare RabbitMQ queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	shared.FailOnError(err, "Failed to register RabbitMQ consumer")

	consumerRoutine := make(chan bool)

	go func() {
		for d := range msgs {
			var reminder models.Reminder

			if err = json.Unmarshal(d.Body, &reminder); err != nil {
				log.Printf("[ERR] Decoding message: %s", err.Error())
				return
			}
			hub.Broadcast <- reminder
		}
	}()

	<-consumerRoutine
}
