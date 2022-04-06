package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatalln("new producer err", err)
	}

	for i := 0; i < 10; i++ {
		err = producer.Publish("demo", []byte(fmt.Sprintf("hello-%d", i)))
		if err != nil {
			log.Fatalln("publish msg err", err)
		}
	}

	producer.Stop()
}
