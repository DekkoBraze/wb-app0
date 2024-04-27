package sender

import (
	"log"

	"github.com/nats-io/nats.go"
)

func StartSender() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Print(err)
	}
	
	nc.Publish("subject", []byte("Hello World"))
}