package receiver

import (
	"context"

	"github.com/ilolss/outbox-pattern/internal/database"
	"github.com/ilolss/outbox-pattern/internal/helpers/errors"
	"github.com/ilolss/outbox-pattern/internal/protogen"
	"github.com/ilolss/outbox-pattern/internal/task"
	"github.com/jackc/pgx/v5"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Receiver struct{}

func (receiver *Receiver) Receive(ctx context.Context, req *protogen.Task) (*emptypb.Empty, error) {
	tx, err := database.DB.BeginTx(context.Background(), pgx.TxOptions{})

	if err != nil {
		errors.Handler(err)
		return &emptypb.Empty{}, err
	}

	defer tx.Rollback(context.Background())

	task := task.Task{Name: req.GetName(), Description: req.GetDescription()}

	_, err = tx.Exec(context.Background(), "INSERT INTO tasks (name, description) VALUES ($1, $2)", task.Name, task.Description)
	if err != nil {
		errors.Handler(err)
		return &emptypb.Empty{}, err
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO outbox (payload, status) VALUES ($1, 'not complited')", task)
	if err != nil {
		errors.Handler(err)
		return &emptypb.Empty{}, err
	}

	tx.Commit(context.Background())
	return &emptypb.Empty{}, nil
}
