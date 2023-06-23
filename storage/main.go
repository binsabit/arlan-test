package main

func main() {
	f := &FileStorage{Dir: "temp"}
	go initStoreConsumer(f)
	go initGetConsumer(f)
	go initGetProducer()

	initStoreProducer()
}
