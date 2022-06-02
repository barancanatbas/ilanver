package config

import "github.com/streadway/amqp"

func Connect() *amqp.Connection {
	con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	return con
}
