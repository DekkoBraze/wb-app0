package main

import (
	"log"
	"fmt"
	
	dbPkg "app0/internal/database"
	"github.com/nats-io/nats.go"
)

func main() {
	db := dbPkg.Database{}
	err := db.Init()
	if (err != nil) {
		log.Print(err)
	}
}