//
//  Practicing RabbitMQ
//
//  Copyright © 2016. All rights reserved.
//

package config

import (
	"fmt"

	"github.com/streadway/amqp"
)

// InitRabbitMQ will create a variable that represent the amqp.Connection
func InitRabbitMQ() (*amqp.Connection, error) {
	client, err := amqp.Dial(Configuration.RabbitMQ.Addr)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping connection to rabbitMQ: %s", err.Error())
	}

	return client, nil
}
