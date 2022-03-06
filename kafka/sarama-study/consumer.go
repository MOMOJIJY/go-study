package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)


func asyncConsume() {
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

	offset := int64(0)
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, offset)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
}


type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error   {

	for t, ps := range sess.Claims() {
		for	_, p := range ps {
			if p == 0 {
				sess.ResetOffset(t, p, 2163044, "")
			} else {
				sess.ResetOffset(t, p, 0, "")
			}
			log.Printf("Reset offset on topic=%s, partition=%d\n", t, p)
		}
	}
	time.Sleep(5*time.Second)
	return nil
}
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Println("partition=", claim.Partition())
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d InitialOffset:%d HighWaterMarkOffset:%d\n",
			msg.Topic, msg.Partition, msg.Offset, claim.InitialOffset(), claim.HighWaterMarkOffset())
		sess.MarkMessage(msg, "")
	}
	return nil
}

func consumeGroup() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0 // specify appropriate version
	config.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup([]string{broker}, group, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{topic}
		handler := exampleConsumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}