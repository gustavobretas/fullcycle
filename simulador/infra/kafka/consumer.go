package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MessageChannel chan *ckafka.Message
}

func NewKafkaConsumer(messageChannel chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChannel: messageChannel,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	consumer.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Fatalf("Error during consuming: %s", err)
		}
		k.MessageChannel <- msg
	}
}
