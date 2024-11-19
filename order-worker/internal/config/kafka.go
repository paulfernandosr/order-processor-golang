package config

import (
	"log"

	"github.com/IBM/sarama"
)

func NewKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer([]string{Props.KafkaBroker}, nil)

	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}

	return consumer
}

func NewKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer([]string{Props.KafkaBroker}, config)

	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	return producer
}
