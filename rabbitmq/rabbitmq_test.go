package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	Queuename string
	Exchange  string
	Key       string
	Mqurl     string
}
