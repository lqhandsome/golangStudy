package main

import (
	"github.com/streadway/amqp"
	"log"
)
func main() {
	//连接服务器
	conn,err := amqp.Dial("amqp://liuqiang:smlaibulai@8.129.122.24:5672/")
	failOnError(err,"发送端连接失败")
	defer conn.Close()

	//创建一个信道
	ch ,err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	args := make(map[string]interface{})
	args["x-message-ttl"] = 5000
	args["x-dead-letter-exchange"] = "dead-exchange"
	//args["x-dead-letter-routing-key"] = "routing" //fonout不需要
	ch.QueueDeclare(
		"ttl_queue", // name
		true,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		args,     // arguments
	)

	err = ch.Publish(
		"",       // exchange
		"", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
		},
	)

	body := "hello"
	err = ch.Publish(
		"",       // exchange
		"ttl_queue", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:"5000",
		})

	failOnError(err,"发送失败")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

