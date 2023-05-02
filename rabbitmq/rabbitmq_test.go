package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}
