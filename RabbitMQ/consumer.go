//Vai consumir as mensagens
package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer starting...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}

	messages, err := ch.Consume(
		"TestQeue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("Recieved message: %s\n", message.Body)
		}
	}()

	fmt.Println("Successfully connected")
	fmt.Println("[*] - Waiting for messages")

	<-forever

	defer ch.Close()
	defer conn.Close()
}
