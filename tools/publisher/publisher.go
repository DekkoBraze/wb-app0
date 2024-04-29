package main

import (
	"io"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

const fileName = "model.json"

// Отправление json на сервер nats-streaming
func main() {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	sc, err := stan.Connect("test-cluster", "test-publisher")
	if err != nil {
		panic(err)
	}

	err = sc.Publish("app0", byteData)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(fileName, " has been sent.")
	}

}
