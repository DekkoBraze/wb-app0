package main

import (
	"log"
	"net/http"

	cachePkg "app0/internal/cache"
	dbPkg "app0/internal/database"
	"app0/internal/handler"
	"app0/internal/subscriber"

	"github.com/gorilla/mux"
)

// Инициализация кэша
var cache = cachePkg.New()

func main() {
	// Инициализия бд
	db := dbPkg.Database{}
	err := db.Init()
	if err != nil {
		panic(err)
	}

	// Восстановление кэша из бд
	err = cache.Recover(db)
	if err != nil {
		panic(err)
	}

	// Подписка и обработка сообщений на nats-streaming
	sc, sub, err := subscriber.MessageProcessing(cache, db)
	if err != nil {
		log.Print(err)
	}
	defer sc.Close()
	defer sub.Close()

	// Сервер
	router := mux.NewRouter()
	router.Handle("/getOrder/{orderId}", &handler.GetOrder{Cache: cache}).Methods("GET", "OPTIONS")
	http.Handle("/", router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
