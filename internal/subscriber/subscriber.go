package subscriber

import (
	"app0/internal/structs"
	"encoding/json"
	"log"

	cachePkg "app0/internal/cache"
	dbPkg "app0/internal/database"

	"github.com/nats-io/stan.go"
)

func MessageProcessing(cache *cachePkg.Cache, db dbPkg.Database) (sc stan.Conn, sub stan.Subscription, err error) {
	// Подключение к серверу nats-streaming
	sc, err = stan.Connect("test-cluster", "test-client")
	if err != nil {
		return
	}

	// Подписка на сабж app0, через который будем передавать сообщения
	sub, err = sc.Subscribe("app0", func(m *stan.Msg) {
		log.Print("The json order has been received...")
		var newOrder structs.Order

		err = json.Unmarshal(m.Data, &newOrder)
		if err != nil {
			log.Print(err)
		}

		// Валидация и сохранение данных при её прохождении
		if structs.OrderValidated(newOrder) {
			err = db.InsertOrder(m.Data, newOrder)
			if err != nil {
				log.Print(err)
				return
			}
			cache.Set(newOrder)
			log.Print("The json order has been cached and added to the db.")
		} else {
			log.Print("Received wrong json order!")
		}
	})
	if err != nil {
		return
	}
	return
}
