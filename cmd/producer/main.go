package main

import (
	"log"

	"github.com/ilolss/outbox-pattern/internal/helpers/errors"
	"github.com/ilolss/outbox-pattern/internal/kafka"
	"github.com/ilolss/outbox-pattern/internal/producer/processing"
	"github.com/ilolss/outbox-pattern/internal/task"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("[WARN] No .env file found")
    }
}

func main() {
	kafka.ReaderConnect()
	defer kafka.ReaderClose()

	for {
		var task task.Task
		err := kafka.Read(&task)
		if err != nil {
			errors.Handler(err)
			continue
		}
		processing.Processing(&task)
	}
}