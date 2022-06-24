package main

import (
	"log"

	"github.com/gustavobretas/fullcycle/simulador/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("Hello Kafka!", "route.new-position", producer)
	for {
		_ = 1
	}
}
