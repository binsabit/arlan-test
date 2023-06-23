package main

import (
	"log"
	"os"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initStoreConsumer(s Storage) {
	// conn
	conn, err := amqp.Dial(rabbitConfig.uri)
	if err != nil {
		log.Printf("ERROR: fail init storage consumer: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("INFO: done init storage consumer conn")

	// create channel
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR: fail create storage channel: %s", err.Error())
		os.Exit(1)
	}

	// create queue
	queue, err := amqpChannel.QueueDeclare(
		"storage", // channelname
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("ERROR: fail create storage queue: %s", err.Error())
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
		log.Printf("ERROR: fail create storage channel: %s", err.Error())
		os.Exit(1)
	}

	// consume
	for {
		select {
		case msg := <-msgChannel:
			// unmarshal
			req := &storage.StoreImageRequest{}
			err = proto.Unmarshal(msg.Body, req)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", msg.Body)
				continue
			}
			log.Printf("INFO: received storage msg: %v", req.Uid)

			// ack for message
			err = msg.Ack(true)
			if err != nil {
				log.Printf("ERROR: fail to ack: %s", err.Error())
			}

			// handle docMsg
			handleStoreRequest(s, req)
		}
	}
}

func handleStoreRequest(s Storage, req *storage.StoreImageRequest) {
	reply := storage.StoreImageResponse{Uid: req.Uid}

	err := s.StoreImage(req)
	if err != nil {
		reply.Status = "Can't store err: " + err.Error()
	} else {
		reply.Status = "Created"
	}

	msg := RabbitStoreMsg{
		QueueName: req.ReplyTo,
		Reply:     storage.StorageRespone{StoreImageResponse: &reply, GetImageResponse: nil},
	}
	log.Println(msg)
	schan <- msg
}
