package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {

	//连接broker
	conn, err := amqp.Dial("amqp://liuqiang:smlaibulai@8.129.122.24:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")
	ch.ExchangeDeclare(
		"direct_order_exchange",
		"direct",
		true,
		false,
		false,
		true,
		nil,
		)

	ch.QueueDeclare(
		"queue5",
		true,
		false,
		false,
		false,
		nil,
		)
	ch.QueueBind(
		"queue5",
		"weixin",
		"direct_order_exchange",
		false,
		nil,
		)
	body := "topic_exchange"
	err = ch.Publish(
		"direct_order_exchange",     // exchange
		"com.order.xxx", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
