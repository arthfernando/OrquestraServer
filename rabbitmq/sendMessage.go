package rabbitmq

import (
	"encoding/json"
	"log"
	"painellembretes/config"
	"painellembretes/models"
	"painellembretes/shared"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage(reminder models.Reminder) error {
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

	encodedBody, err := json.Marshal(reminder)
	if err != nil {
		log.Println("Failed to encode body")
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(encodedBody),
		},
	)
	shared.FailOnError(err, "Failed to publish message")
	return nil
}
