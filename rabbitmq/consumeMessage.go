package rabbitmq

import (
	"encoding/json"
	"log"
	"painellembretes/config"
	"painellembretes/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

// func ConsumeMessage() (<-chan amqp.Delivery, error) {
func ConsumeMessage(bodyMessage chan models.Reminder) {
	rabbitmqUrl := config.Get("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitmqUrl)
	// shared.FailOnError(err, "Failed to connect to RabbitMQ")
	// if err != nil {
	// 	return nil, errors.New("failed to connect to RabbitMQ")
	// }
	defer conn.Close()

	ch, err := conn.Channel()
	// shared.FailOnError(err, "Failed to open RabbitMQ channel")
	// if err != nil {
	// 	return nil, errors.New("failed to open RabbitMQ channel")
	// }
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"painellembretes-server",
		false,
		false,
		false,
		false,
		nil,
	)
	// shared.FailOnError(err, "Failed to declare RabbitMQ queue")
	// if err != nil {
	// 	return nil, errors.New("failed to declare RabbitMQ queue")
	// }

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	// shared.FailOnError(err, "Failed to register RabbitMQ consumer")
	// if err != nil {
	// 	return nil, errors.New("failed to register RabbitMQ consumer")
	// }

	consumerRoutine := make(chan bool)

	// return msgs, nil

	go func() {
		for d := range msgs {
			var reminder models.Reminder

			if err = json.Unmarshal(d.Body, &reminder); err != nil {
				log.Printf("[ERR] Decoding message: %s", err.Error())
				return
			}
			log.Printf("Received message: %s", reminder)
			bodyMessage <- reminder
		}
	}()

	log.Println("[*] Waiting messages. Press CTRL+C to exit.")
	<-consumerRoutine
}
