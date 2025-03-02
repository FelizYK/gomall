package mq

import (
	"log"

	"github.com/nats-io/nats.go"
)

var (
	nc *nats.Conn
)

func Init() {
	var err error
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect nats: %v", err)
	}

	initSubscriber()
}

func Close() {
	if nc != nil {
		closeSubscriber()
		nc.Close()
	}
}
