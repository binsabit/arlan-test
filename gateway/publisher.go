package main

import (
	"log"
	"os"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitStoreMsg struct {
	QueueName string
	Message   storage.StoreImageRequest
}

type RabbitGetImgMsg struct {
	QueueName string
	Message   storage.GetImageRequest
}

// channel to publish rabbit messages
var schan = make(chan RabbitStoreMsg, 10)
var gchan = make(chan RabbitGetImgMsg, 10)

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
		case msg := <-schan:
			// marshal
			data, err := proto.Marshal(&msg.Message)
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

			log.Printf("INFO: published to \"storage\"msg: %v", &msg.Message.Uid)

		case msg := <-gchan:
			// marshal
			data, err := proto.Marshal(&msg.Message)
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

			log.Printf("INFO: published to \"image\" msg: %v", &msg.Message.Name)
		}
	}
}
