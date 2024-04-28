package main

import (
	"encoding/json"
	"log"
	"net/http"

	cachePkg "app0/internal/cache"
	dbPkg "app0/internal/database"

	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
)

var cache = cachePkg.New()

func main() {
	db := dbPkg.Database{}
	err := db.Init()
	if err != nil {
		log.Print(err)
	}

	var order dbPkg.Order
	dbOrders, err := db.SelectOrders()
	if err != nil {
		log.Print(err)
	}
	for _, byteOrder := range dbOrders {
		json.Unmarshal(byteOrder, &order)
		cache.Set(order)
	}

	//cache.PrintItems()

	sc, err := stan.Connect("test-cluster", "test-client")
	if err != nil {
		log.Print(err)
	}

	_, err = sc.Subscribe("app0", func(m *stan.Msg) {
		err = json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Print(err)
		} 
		log.Print(order)
		cache.Set(order)
		err := db.InsertJson(m.Data)
		if err != nil {
			log.Print(err)
		}
	})
	if err != nil {
		log.Print(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/getOrder/{orderId}", get_order).Methods("GET", "OPTIONS")
	http.Handle("/", router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func get_order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	order_uid := vars["orderId"]
	sendedOrder, orderIsFound := cache.Get(order_uid)
	if !orderIsFound {
		log.Print("cant find any orders!")
		return
	}
	json.NewEncoder(w).Encode(sendedOrder)
}
