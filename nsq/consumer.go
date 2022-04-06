package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type myMsgHandler struct {
}

func (myMsgHandler) HandleMessage(message *nsq.Message) error {
	msgId := string(message.ID[:])
	msgBody := string(message.Body)
	log.Printf("hand message msgId:%v msgBody:%v", msgId, msgBody)
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("demo", "channel01", config)
	if err != nil {
		log.Fatalln(err)
	}
	consumer.AddHandler(myMsgHandler{})
	consumer.ConnectToNSQD("127.0.0.1:4150")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumer.Stop()
}
