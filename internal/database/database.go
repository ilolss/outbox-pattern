package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Open(connString string) {
	connConf, err := pgxpool.ParseConfig(connString)
    if err != nil {
		// TODO: think about it
		log.Fatalln("[ERROR] ", err.Error())
	}
    connConf.MaxConns = 50
	DB, err = pgxpool.NewWithConfig(context.Background(), connConf)
	if err != nil {
		// TODO: think about it
		log.Fatalln("[ERROR] ", err.Error())
	}
}

func Close() {
	DB.Close()
}
