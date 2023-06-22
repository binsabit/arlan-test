package main

import "log"

func main() {

	go initConsumer()

	go initProducer()

	gatewayApi := NewGatewayAPI(":4000")

	err := gatewayApi.runApi()
	if err != nil {
		log.Fatal(err)
	}

}
