package main

import (
	"io"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func main() {
	jsonFile, err := os.Open("model.json")
	if err != nil {
		log.Print(err)
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Print(err)
	}

	sc, err := stan.Connect("test-cluster", "test-publisher")
	if err != nil {
		log.Print(err)
	}

	err = sc.Publish("app0", byteData)
	if err != nil {
		log.Print(err)
	}

}
