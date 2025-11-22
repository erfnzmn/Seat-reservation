package rabbitmq

import (
	"github.com/streadway/amqp"
)

func ConsumeMessages(ch *amqp.Channel, queue string) (<-chan amqp.Delivery, error) {
	_, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}
