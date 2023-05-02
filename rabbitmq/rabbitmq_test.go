package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}
