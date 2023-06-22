package main

import (
	"log"
	"os"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMsg struct {
	QueueName string                     `json:"queueName"`
	Reply     storage.StoreImageResponse `json:"reply"`
}

// channel to publish rabbit messages
var rchan = make(chan RabbitMsg, 10)

func initProducer() {
	// conn
	conn, err := amqp.Dial(rabbitConfig.uri)
	if err != nil {
		log.Printf("ERROR: fail init consumer: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: done init producer conn")

	// create channel
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR: fail create channel: %s", err.Error())
		os.Exit(1)
	}

	for {
		select {
		case msg := <-rchan:
			// marshal
			data, err := proto.Marshal(&msg.Reply)
			if err != nil {
				log.Printf("ERROR: fail marshal: %s", err.Error())
				continue
			}

			// publish message
			err = amqpChannel.Publish(
				"",            // exchange
				msg.QueueName, // routing key
				false,         // mandatory
				false,         // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        data,
				},
			)
			if err != nil {
				log.Printf("ERROR: fail publish msg: %s", err.Error())
				continue
			}

			log.Printf("INFO: published msg: %v to: %s", &msg.Reply, msg.QueueName)
		}
	}
}