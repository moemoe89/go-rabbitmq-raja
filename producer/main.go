//
//  Practicing RabbitMQ
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	conf "github.com/moemoe89/go-rabbitmq-raja/producer/config"

	"encoding/json"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

type SM struct {
	Hello string `json:"string"`
}

// server struct represent the delivery for controller
type ctrl struct{
	client *amqp.Connection
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(client *amqp.Connection) *ctrl {
	return &ctrl{client}
}

func (c *ctrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ch, err := c.client.Channel()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "`+err.Error()+`"}`))
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "`+err.Error()+`"}`))
		return
	}

	var sm SM
	sm.Hello = "world"

	body, err := json.Marshal(sm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "`+err.Error()+`"}`))
		return
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Success sent message"}`))
	return
}

func main() {
	conn, err := conf.InitRabbitMQ()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctrl := NewCtrl(conn)

	http.HandleFunc("/", ctrl.ServeHTTP)
	log.Fatal(http.ListenAndServe(":"+conf.Configuration.Port, nil))
}
