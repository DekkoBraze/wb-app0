package main

import (
	"log"
	_ "fmt"
	"os"
	"io"
	"encoding/json"

	dbPkg "app0/internal/database"
	"github.com/nats-io/stan.go"
)

func main() {
	db := dbPkg.Database{}
	err := db.Init()
	if (err != nil) {
		log.Print(err)
	}

	sc, err := stan.Connect("test-cluster", "test-client")
	if err != nil {
		log.Print(err)
	}

	var orders dbPkg.Orders

	sub, err := sc.Subscribe("app0", func(m *stan.Msg) {
		json.Unmarshal(m.Data, &orders)
		err := db.InsertJson(m.Data)
		if err != nil {
			log.Print(err)
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Print(err)
	}

	jsonFile, err := os.Open("model.json")
	if err != nil {
		log.Print(err)
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Print(err)
	}

	err = sc.Publish("app0", byteData)
	if err != nil {
		log.Print(err)
	}
	
	sub.Unsubscribe()
}
