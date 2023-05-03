package main

import (
	"fmt"
	"rabgobitmq_go/Rabbitmq"
)

func main() {
	simple := Rabbitmq.NewRabbitMQSimple("ruoye")
	simple.PublishSimple("hello rabbitmq")
	fmt.Println("send success")
}
