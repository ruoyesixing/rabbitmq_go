package main

import "rabgobitmq_go/Rabbitmq"

func main() {
	simple := Rabbitmq.NewRabbitMQSimple("ruoye")
	simple.ConsumeSimple()
}
