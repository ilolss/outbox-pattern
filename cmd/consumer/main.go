package main

import (
	"log"
	"net"
	"os"
	"flag"

	"github.com/ilolss/outbox-pattern/internal/consumer/receiver"
	"github.com/ilolss/outbox-pattern/internal/consumer/sheduler"
	"github.com/ilolss/outbox-pattern/internal/database"
	"github.com/ilolss/outbox-pattern/internal/kafka"
	"github.com/ilolss/outbox-pattern/internal/protogen"
	"github.com/joho/godotenv"

	grpc "google.golang.org/grpc"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("[WARN] No .env file found")
    }
}

func main() {
	database.Open(os.Getenv("PGSQL_DSN"))
	defer database.Close()

	kafka.WriterConnect()
	defer kafka.WriterClose()

	sheduler := sheduler.Start()
	defer close(sheduler)

	host := flag.String("host", "127.0.0.1", "The host to connect to")
	port := flag.String("port", "8080", "The port to connect to")

	flag.Parse()

	s := grpc.NewServer()
	srv := &receiver.Receiver{}
	protogen.RegisterReceiverServer(s, srv)

	l, err := net.Listen("tcp", *host + ":" + *port)
	if err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}

	// TODO: Server
	if err := s.Serve(l); err != nil {
		log.Fatalln("[ERROR] ", err.Error())
	}
}