package repository

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/config"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/model"
)

type OrderProducer interface {
	Produce(order *model.Order)
}

type KafkaOrderProducer struct {
	syncProducer sarama.SyncProducer
}

func NewKafkaOrderProducer(producer sarama.SyncProducer) OrderProducer {
	return &KafkaOrderProducer{producer}
}

func (orderProducer *KafkaOrderProducer) Produce(order *model.Order) {
	messageBytes, err := json.Marshal(order)

	if err != nil {
		log.Fatalf("Error serializing the message: %v", err)
	}

	message := &sarama.ProducerMessage{
		Topic: config.Props.OrderTopic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := orderProducer.syncProducer.SendMessage(message)

	if err != nil {
		log.Fatalf("Error sending message to kafka: %v", err)
	}

	log.Printf("Message sent to the partition %d, offset %d\n", partition, offset)
}
