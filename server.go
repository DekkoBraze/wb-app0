package main

import (
	"encoding/json"
	"log"
	"net/http"

	cachePkg "app0/internal/cache"
	dbPkg "app0/internal/database"
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
	router.HandleFunc("/getOrder/{orderId}", getOrder).Methods("GET", "OPTIONS")
	http.Handle("/", router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

// Выдача json заказа
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Print("Received json order request...")
	vars := mux.Vars(r)
	order_uid := vars["orderId"]
	sendedOrder, orderIsFound := cache.Get(order_uid)
	if !orderIsFound {
		log.Print("No orders were found!")
		return
	}
	json.NewEncoder(w).Encode(sendedOrder)
	log.Print("The json order has been sent.")
}
