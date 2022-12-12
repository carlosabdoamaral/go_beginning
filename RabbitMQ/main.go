package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	// Conecta no rabbitmq usando o amqp
	// amqp://username:password@url
	fmt.Println("Connecting to RabbitMQ...")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	fmt.Println("Connected!")

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	// Criar uma qeue
	q, err := ch.QueueDeclare(
		"TestQeue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(q)

	// Manda a mensagem do cliente para o exchange e entra na fila do RabbitMQ
	err = ch.Publish(
		"",
		"TestQeue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello RabbitMQ!"),
		},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Successfully published message to the qeue")
}
