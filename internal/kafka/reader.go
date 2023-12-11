package kafka

import (
	"context"
	"log"
	"os"

	"github.com/ilolss/outbox-pattern/internal/task"
	"github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func ReaderConnect() {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		Topic:   os.Getenv("KAFKA_TOPIC_NAME"),
	})
}

func Read(task *task.Task) error {
	message, err := reader.ReadMessage(context.Background())

	if err != nil {
		return err
	}

	err = task.FromJSON(message.Value)
	return err
}

func ReaderClose() {
	err := reader.Close()
	if err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}
}