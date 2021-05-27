package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
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

	body := "hello"
	for i := 0; i < 20; i++ {
		//time.Sleep(time.Second/2)
		err = ch.Publish(
			"",       // exchange
			"queue1", // routing key
			false,    // mandatory
			false,    // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body+strconv.Itoa(i)),
			})
		failOnError(err, "Failed to publish a message")

	}

}
