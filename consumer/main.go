//
//  Practicing RabbitMQ
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	conf "github.com/moemoe89/go-rabbitmq-raja/consumer/config"

	"encoding/json"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type SM struct {
	Hello string `json:"string"`
}

func main() {
	conn, err := conf.InitRabbitMQ()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			text := string(d.Body)
			bytes := []byte(text)

			var sm SM
			json.Unmarshal(bytes, &sm)

			log.Printf("Received a message: %s", sm.Hello)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}