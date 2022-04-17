package queue

import (
	"encoding/json"
	"fmt"
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"strconv"

	"ilanver/pkg/logger"

	"github.com/streadway/amqp"
)

type Queue struct {
	connection *amqp.Connection
	queue      amqp.Queue
	channel    *amqp.Channel
}

func NewQueue() *Queue {

	connection := config.Connect()

	qu := &Queue{
		connection: connection,
	}

	return qu
}

func (q *Queue) Publish(queueName string, message []byte) error {

	defer q.connection.Close()
	channel, err := q.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	_, err = declareQueue(channel, queueName)
	if err != nil {
		fmt.Println("error declaring queue: ", err)
		return err
	}

	err = channel.Publish(
		"", // exchange
		queueName,
		false, // mandatory
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("success publish message to queue")
	return nil
}

func ConsumeInsertProduct(queueName string) {

	connection := config.Connect()
	defer connection.Close()
	channel, err := connection.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		panic(err)
	}

	msg, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if err != nil {
		panic(err)
	}

	repo := repository.NewProductElasticRepository(config.ElasticDB)

	forever := make(chan bool)
	go func() {
		for d := range msg {
			product := model.ProductElastic{}
			json.Unmarshal(d.Body, &product)

			id := strconv.Itoa(int(product.ID))

			err := repo.Save(d.Body, id)
			if err != nil {
				logger.Errorf(4, "error save product to elastic: %v", err)
			}
		}
	}()

	fmt.Println("listening message for insert product ...")
	<-forever
}

func declareQueue(channel *amqp.Channel, queueName string) (amqp.Queue, error) {
	q, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println("error declaring queue: ", err)
		return q, err
	}
	return q, nil
}
