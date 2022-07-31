package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

func idempotence() {
	logger := &log.Logger{}
	logger.SetOutput(os.Stdout)
	sarama.Logger = logger

	config := sarama.NewConfig()
	config.Producer.Idempotent = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Net.MaxOpenRequests = 1

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "topic-test"

	for i := 0; i <= 10; i++ {
		text := fmt.Sprintf("message %08d", i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(text),
		}
		producer.Input() <- msg
	}

	time.Sleep(5 * time.Second)
}