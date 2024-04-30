package handler

import (
	cachePkg "app0/internal/cache"
	"app0/internal/structs"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Ответ сервера на запрос
type Response struct {
	Message string        `json:"message"`
	Order   structs.Order `json:"order"`
}

// Хендлер
type GetOrder struct {
	Cache *cachePkg.Cache
}

// Выдача json заказа
func (h *GetOrder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	log.Print("Received json order request...")
	vars := mux.Vars(r)
	order_uid := vars["orderId"]
	sendedOrder, orderIsFound := h.Cache.Get(order_uid)
	if !orderIsFound {
		var response = Response{Message: "ORDERS_NOT_FOUND", Order: structs.Order{}}
		json.NewEncoder(w).Encode(response)
		log.Print("No orders were found!")
		return
	}

	var response = Response{Message: "OK", Order: sendedOrder}
	json.NewEncoder(w).Encode(response)
	log.Print("The json order has been sent.")
}
