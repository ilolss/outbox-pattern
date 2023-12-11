package sheduler

import (
	"context"
	"log"
	"time"

	"github.com/ilolss/outbox-pattern/internal/database"
	"github.com/ilolss/outbox-pattern/internal/helpers/errors"
	"github.com/ilolss/outbox-pattern/internal/kafka"
	"github.com/ilolss/outbox-pattern/internal/task"

	"github.com/jackc/pgx/v5"
)

// стартует sheduler, то есть каждую секунду обращается в служебной бд и вытаскивает новое задание,
// затем отправляет его в kafka, если все хорошо, то помечает его, как сделанное
func Start() chan struct{} {
	ticker := time.NewTicker(time.Second * 5)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			 case <- ticker.C:
				tx, err := database.DB.BeginTx(context.Background(), pgx.TxOptions{})
				if err != nil {
					errors.Handler(err)
					continue
				}
				defer tx.Rollback(context.Background())

				var id int
				var task task.Task;

				err = tx.QueryRow(context.Background(), "SELECT id, payload FROM outbox WHERE status = 'not complited' ORDER BY id LIMIT 1 FOR UPDATE SKIP LOCKED").Scan(&id, &task)

				if err != nil {
					errors.Handler(err)
					tx.Rollback(context.Background())
					continue
				}

				log.Println("[INFO] ", task)
				err = kafka.Write(task)
				if err != nil {
					errors.Handler(err)
					tx.Rollback(context.Background())
					continue
				}

				_, err = tx.Exec(context.Background(), "UPDATE outbox SET status = 'complited' WHERE id = $1", id)

				if err != nil {
					errors.Handler(err)
					tx.Rollback(context.Background())
					continue
				}

				tx.Commit(context.Background())
			 case <- quit:
				ticker.Stop()
				return
			 }
		 }
	}()
	return quit
}