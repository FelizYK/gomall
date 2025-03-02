package mq

import (
	"log"

	"github.com/FelizYK/gomall/app/email/service"
	rpcemail "github.com/FelizYK/gomall/rpc/email"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

var (
	sub *nats.Subscription
)

func initSubscriber() {
	var err error
	sub, err = nc.Subscribe("email", func(m *nats.Msg) {
		var req rpcemail.SendEmailReq
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			log.Printf("proto.Unmarshal error: %v\n", err)
			return
		}
		err = service.SendEmail(&req)
		if err != nil {
			log.Printf("service.SendEmail error: %v\n", err)
			return
		}
	})
	if err != nil {
		log.Fatalf("nc.Subscribe error: %v\n", err)
	}
}

func closeSubscriber() {
	sub.Unsubscribe()
}
