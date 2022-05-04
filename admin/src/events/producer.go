package events

import (
	"encoding/json"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer *kafka.Producer

func SetupProducer() {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("SECURITY_PROTOCOL"),
		"sasl.username":     os.Getenv("SASL_USERNAME"),
		"sasl.password":     os.Getenv("SASL_PASSWORD"),
		"sasl.mechanism":    os.Getenv("SASL_MECHANISM"),
	})
	if err != nil {
		panic(err)
	}

	// defer Producer.Close()
}

func Produce(topic, key string, message interface{}) {
	value, _ := json.Marshal(message)

	Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: int32(kafka.PartitionAny),
		},
		Key:   []byte(key),
		Value: value,
	}, nil)

	Producer.Flush(15000)
}
