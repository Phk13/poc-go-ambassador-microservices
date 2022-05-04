package main

import (
	"checkout/src/database"
	"checkout/src/events"
	"checkout/src/models"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	database.Connect()

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
		"sasl.mechanism":    os.Getenv("SASL_MECHANISM"),
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC")}, nil)

	defer consumer.Close()

	fmt.Println("Started consuming")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			database.DB.Create(&models.KafkaError{
				Key:   msg.Key,
				Value: msg.Value,
				Error: err,
			})
			return
		}
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		if err := events.Listen(msg); err != nil {
			database.DB.Create(&models.KafkaError{
				Key:   msg.Key,
				Value: msg.Value,
				Error: err,
			})
		}
	}
}
