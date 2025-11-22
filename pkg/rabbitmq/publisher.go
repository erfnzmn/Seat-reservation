package rabbitmq

import (
	"github.com/streadway/amqp"
)

func PublishMessage(ch *amqp.Channel, queue string, body []byte) error {
	_, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
