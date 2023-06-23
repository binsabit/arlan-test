package main

import (
	"log"
	"os"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initConsumer() {
	// conn
	conn, err := amqp.Dial(rabbitConfig.uri)
	if err != nil {
		log.Printf("ERROR: fail init consumer: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: done init consumer conn")

	// create channel
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR: fail create channel: %s", err.Error())
		os.Exit(1)
	}

	// create queue
	queue, err := amqpChannel.QueueDeclare(
		"gateway", // channelname
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("ERROR: fail create queue: %s", err.Error())
		os.Exit(1)
	}

	// channel
	msgChannel, err := amqpChannel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Printf("ERROR: fail create channel: %s", err.Error())
		os.Exit(1)
	}

	// consume
	for {
		select {
		case msg := <-msgChannel:
			// unmarshal
			storageResponse := storage.StorageRespone{}
			err = proto.Unmarshal(msg.Body, &storageResponse)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", msg.Body)
				continue
			}

			getImageResponse := storageResponse.GetImageResponse
			storeImageResponse := storageResponse.StoreImageResponse

			if getImageResponse != nil {
				log.Printf("INFO: get image received msg status: %v", getImageResponse.Status)

				// ack for message
				err = msg.Ack(true)
				if err != nil {
					log.Printf("ERROR: fail to ack: %s", err.Error())
				}

				// find waiting channel(with uid) and forward the reply to it
				if grchan, ok := grchans[getImageResponse.Image.Name]; ok {
					grchan <- *getImageResponse
				}
			} else if storeImageResponse != nil {
				log.Printf("INFO: store img received msg status : %v", storeImageResponse.Status)

				// ack for message
				err = msg.Ack(true)
				if err != nil {
					log.Printf("ERROR: fail to ack: %s", err.Error())
				}

				// find waiting channel(with uid) and forward the reply to it
				if srchan, ok := srchans[storeImageResponse.Uid]; ok {
					srchan <- *storeImageResponse
				}
			}

		}
	}
}
