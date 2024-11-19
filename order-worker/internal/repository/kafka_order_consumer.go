package repository

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/config"
	"github.com/paulfernandosr/order-processor-golang/order-worker/internal/model"
)

type OrderConsumer interface {
	Consume(orderChan chan *model.Order)
}

type KafkaOrderConsumer struct {
	partitionConsumer sarama.PartitionConsumer
}

func NewKafkaOrderConsumer(consumer sarama.Consumer) OrderConsumer {
	partitionConsumer, err := consumer.ConsumePartition(config.Props.OrderTopic, 0, sarama.OffsetNewest)

	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}

	return &KafkaOrderConsumer{partitionConsumer}
}

func (orderConsumer *KafkaOrderConsumer) Consume(orderChan chan *model.Order) {
	defer orderConsumer.partitionConsumer.Close()

	for {
		message, ok := <-orderConsumer.partitionConsumer.Messages()

		if !ok {
			log.Println("Partition consumer channel closed")
			break
		}

		var order model.Order

		err := json.Unmarshal(message.Value, &order)

		if err != nil {
			log.Println("Error unmarshalling order")
			continue
		}

		log.Printf("Order message received successfully: %+v\n", order)

		orderChan <- &order
	}

	close(orderChan)
}
