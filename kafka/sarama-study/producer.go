package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"myProject/go-study/kafka/sarama-study/interceptor"
	"myProject/go-study/kafka/sarama-study/partitioner"
)

func asyncProduce() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = partitioner.NewMyPartitioner
	config.Producer.Interceptors = []sarama.ProducerInterceptor{interceptor.NewMyInterceptor()}

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder("testing 123"), Partition: 2}
		select {
		case producer.Input() <- message:
			enqueued++
			time.Sleep(3 * time.Second)
			log.Printf("Successfully sended %d msg\n", enqueued)

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)
}


