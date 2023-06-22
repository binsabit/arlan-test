package main

import (
	"log"
	"os"
	"path"

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
		"storage", // channelname
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
			docMsg := &storage.StoreImageRequest{}
			err = proto.Unmarshal(msg.Body, docMsg)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", msg.Body)
				continue
			}
			log.Printf("INFO: received msg: %v", docMsg.Uid)

			// ack for message
			err = msg.Ack(true)
			if err != nil {
				log.Printf("ERROR: fail to ack: %s", err.Error())
			}

			// handle docMsg
			handleMsg(docMsg)
		}
	}
}

func handleMsg(docMsg *storage.StoreImageRequest) {
	// TODO create doc on storage
	// log.Println(docMsg.Image.Name)
	filepath := path.Join("./temp", docMsg.Image.Name)
	_ = os.Mkdir(filepath, os.ModePerm)
	file, err := os.OpenFile(path.Join(filepath, docMsg.Image.Name+".webp"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {

	}
	file.Write([]byte(docMsg.Image.Content))
	defer file.Close()

	reply := storage.StoreImageResponse{
		Uid:    docMsg.Uid,
		Status: "Created",
	}
	msg := RabbitMsg{
		QueueName: docMsg.ReplyTo,
		Reply:     reply,
	}
	rchan <- msg
}
