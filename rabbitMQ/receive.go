package rsrabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

// func failOnErr(err error, msg string) {
// 	if err != nil {
// 		fmt.Printf("%s: %s", msg, err)
// 	}
// }

// Receive test amqp receive func
func Receive() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErr(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnErr(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnErr(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s \n", d.Body)
		}
	}()

	fmt.Printf(" [*] Waiting for message. To exit press CTRL+C \n")
	<-forever
}
