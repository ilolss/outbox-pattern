package kafka

import (
	"context"
	"log"
	"os"

	"github.com/ilolss/outbox-pattern/internal/task"
	"github.com/segmentio/kafka-go"
)

var writer kafka.Writer

func WriterConnect() {
	writer = kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_HOST")),
		Topic: os.Getenv("KAFKA_TOPIC_NAME"),
	}
}

func Write(task task.Task) error {
	payload, err := task.ToJSON()

	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{Value: payload})

	return err
}

func WriterClose() {
	err := writer.Close()
	if err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}
}
