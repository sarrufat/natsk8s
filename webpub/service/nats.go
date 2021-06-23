package service

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type (
	NatsClient struct {
		url string
	}
	NatsAction func(conn *nats.Conn)
)

func (cli *NatsClient) Do(action NatsAction) {
	nc, err := nats.Connect(cli.url, nats.Name("webpub"), nats.Timeout(5*time.Second))
	defer nc.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		action(nc)
	}
}
