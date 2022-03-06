package partitioner

import (
	"log"

	"github.com/Shopify/sarama"
)

type myPartitioner struct {
	partition int32
}

func (p *myPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	if p.partition >= numPartitions {
		p.partition = 0
	}
	ret := p.partition
	p.partition++
	log.Printf("[Partitioner] partition: %d\n", p.partition)
	return ret, nil
}

func (p *myPartitioner) RequiresConsistency() bool {
	return false
}

func NewMyPartitioner(topic string) sarama.Partitioner {
	return &myPartitioner{}
}