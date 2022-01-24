package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

const (
	topic = "demo-topic"
	broker = "localhost:9092"
	group = "my-group"
)

func main() {
	mode := flag.String("mode", "producer", "kafka mode")
	flag.Parse()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	switch *mode {
	case "producer":
		go asyncProduce()
	case "consumer":
		//asyncConsume()
		go consumeGroup()
	case "metadata":
		metadata()
	default:
		panic("invalid mode")
	}

	<-ch
}

