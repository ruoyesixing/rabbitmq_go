package Rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		log.Panicf(fmt.Sprintf("%s:%s", message, err))
	}
}

func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.Conn.Close()
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnErr(err, "failed to connect Rabbitmq")
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	rabbitMQ.failOnErr(err, "failed to open a channel")
	return rabbitMQ
}

// 生产者
func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.Channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	r.Channel.Publish(r.Exchange, r.QueueName, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 消费者
func (r *RabbitMQ) ConsumeSimple() {
	queue, err := r.Channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	message, err := r.Channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for delivery := range message {
			log.Printf("Receive a message: %s", delivery.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
