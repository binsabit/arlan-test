package main

import (
	"log"
	"os"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initGetConsumer(s Storage) {
	// conn
	conn, err := amqp.Dial(rabbitConfig.uri)
	if err != nil {
		log.Printf("ERROR: fail init image consumer: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: done init get image consumer conn")

	// create channel
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR: fail create channel for image: %s", err.Error())
		os.Exit(1)
	}

	// create queue
	queue, err := amqpChannel.QueueDeclare(
		"image", // channelname
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("ERROR: fail create queue for image: %s", err.Error())
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
		log.Printf("ERROR: fail create channel for image: %s", err.Error())
		os.Exit(1)
	}

	// consume
	for {
		select {
		case msg := <-msgChannel:
			// unmarshal
			req := &storage.GetImageRequest{}
			err = proto.Unmarshal(msg.Body, req)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", msg.Body)
				continue
			}
			log.Printf("INFO: received get image msg for: %v", req.Name)

			// ack for message
			err = msg.Ack(true)
			if err != nil {
				log.Printf("ERROR: fail to ack: %s", err.Error())
			}

			// handle docMsg
			handleGetRequest(s, req)
		}
	}
}

func handleGetRequest(s Storage, req *storage.GetImageRequest) {

	reply := storage.GetImageResponse{}
	img, err := s.GetImage(req)
	if err != nil {
		reply.Status = "Not Found, error: " + err.Error()
	} else {
		reply.Status = "Success"
		reply.Image = img
	}

	msg := RabbitGetMsg{
		QueueName: req.ReplyTo,
		Reply:     storage.StorageRespone{GetImageResponse: &reply, StoreImageResponse: nil},
	}
	gchan <- msg
}
