package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func metadata() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	topics, err := consumer.Topics()
	if err != nil {
		log.Fatalln(err)
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("topics=%v\n", topics)
	log.Printf("partitons=%v\n", partitions)
}
