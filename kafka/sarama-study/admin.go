package main

import (
	"log"

	"github.com/Shopify/sarama"
)

type Admin struct {
	client sarama.ClusterAdmin
}

func NewAdmin() {
	config := sarama.NewConfig()
	client, err := sarama.NewClusterAdmin([]string{broker}, config)
	if err != nil {
		panic(err)
	}

	a := &Admin{client: client}

	switch action {
	case "list":
		a.GetTopics()
	case "describe":
		a.DescribeTopic(ktopic)
	case "alter":
		a.AlterTopic(ktopic)
	default:
		panic("invalid action")
	}
}

func (a *Admin) GetTopics() {
	ret, err := a.client.ListTopics()
	if err != nil {
		panic(err)
	}

	for topicName, topicDetail := range ret {
		log.Printf("topicName=%s, NumPartitions=%d, ReplicationFactor=%d, ReplicaAssignment=%v, ConfigEntries=%v\n",
			topicName, topicDetail.NumPartitions, topicDetail.ReplicationFactor, topicDetail.ReplicaAssignment, topicDetail.ConfigEntries)
	}
}

func (a *Admin) DescribeTopic(topic string) {
	if topic == "" {
		panic("empty topic")
	}
	ret, err := a.client.DescribeTopics([]string{topic})
	if err != nil {
		panic(err)
	}

	if len(ret) <= 0 {
		panic("invalid topic")
	}
	meta := ret[0]
	log.Printf("topicName=%s, NumPartitions=%d\n", meta.Name, len(meta.Partitions))
	for _, par := range meta.Partitions {
		log.Printf("partitionID=%d, replicas=%v\n", par.ID, par.Replicas)
	}
}

func (a *Admin) AlterTopic(topic string) {
	if topic == "" {
		panic("empty topic")
	}
}