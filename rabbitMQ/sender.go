package rsrabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Send test amqp connect
func Send(msg string) {
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

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	fmt.Printf(" [x] Sent %s", body)
	failOnErr(err, "Failed to publish a message")

}

func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}
}
