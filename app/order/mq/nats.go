package mq

import (
	"log"

	"github.com/nats-io/nats.go"
)

var (
	Nc *nats.Conn
)

func Init() {
	var err error
	Nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect nats: %v", err)
	}
}

func Close() {
	if Nc != nil {
		Nc.Close()
	}
}
