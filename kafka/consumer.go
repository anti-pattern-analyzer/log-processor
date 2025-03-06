package kafka

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
	"log-processor/services"
	"log-processor/util"
)

func InitializeKafkaReader() *kafka.Reader {
	kafkaURL := getEnv("KAFKA_BROKERS", "localhost:29091,localhost:29092")
	topic := getEnv("KAFKA_TOPIC", "logs-topic")
	groupID := getEnv("KAFKA_GROUP", "log-processor-group")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        strings.Split(kafkaURL, ","),
		Topic:          topic,
		GroupID:        groupID,
		MaxBytes:       10e6,
		MinBytes:       1e3,
		CommitInterval: 0,
		StartOffset:    kafka.FirstOffset,
	})

	log.Println("Kafka Reader Initialized - Listening on", topic)
	return reader
}

func ConsumeKafkaMessages(reader *kafka.Reader) {
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		logLine := string(message.Value)
		logDTO, err := util.ParseLogLine(logLine)
		if err != nil {
			log.Printf("Error parsing log line: %v", err)
			continue
		}

		log.Printf("Parsed LogDTO: %+v\n", logDTO)
		err = services.ProcessRowLog(logDTO)
		if err != nil {
			log.Printf("Error processing log: %v", err)
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
