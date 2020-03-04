//
//  Practicing RabbitMQ
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	conf "github.com/moemoe89/practicing-rabbitmq-golang/producer/config"

	"encoding/json"
	"log"

	"github.com/streadway/amqp"

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

	var sm SM

	sm.Hello = "world"

	body, err := json.Marshal(sm)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
