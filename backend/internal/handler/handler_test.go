package handler

import (
	cachePkg "app0/internal/cache"
	"app0/internal/structs"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetOrderExisting(t *testing.T) {
	// Создание записи в кэше
	cache := cachePkg.New()
	jsonFile, err := os.Open("../../tools/publisher/model.json")
	if err != nil {
		t.Fatal(err)
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Fatal(err)
	}

	var sendingOrder structs.Order

	json.Unmarshal(byteData, &sendingOrder)
	cache.Set(sendingOrder)

	// Отправка реквеста
	req, err := http.NewRequest("GET", "/getOrder/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"orderId": "b563feb7b2b84b6test"})

	rr := httptest.NewRecorder()
	handler := GetOrder{Cache: cache}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response

	json.Unmarshal(rr.Body.Bytes(), &response)

	rightMessage := "OK"
	if response.Message != rightMessage {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, rightMessage)
	}
}

func TestGetOrderNotExisting(t *testing.T) {
	// Создание записи в кэше
	cache := cachePkg.New()
	jsonFile, err := os.Open("../../tools/publisher/model.json")
	if err != nil {
		t.Fatal(err)
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Fatal(err)
	}

	var sendingOrder structs.Order

	json.Unmarshal(byteData, &sendingOrder)
	cache.Set(sendingOrder)

	// Отправка реквеста
	req, err := http.NewRequest("GET", "/getOrder/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"orderId": "wrong_id"})

	rr := httptest.NewRecorder()
	handler := GetOrder{Cache: cache}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response

	json.Unmarshal(rr.Body.Bytes(), &response)

	rightMessage := "ORDERS_NOT_FOUND"
	if response.Message != rightMessage {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, rightMessage)
	}
}

func TestGetOrderEmpty(t *testing.T) {
	// Создание записи в кэше
	cache := cachePkg.New()
	jsonFile, err := os.Open("../../tools/publisher/model.json")
	if err != nil {
		t.Fatal(err)
	}
	defer jsonFile.Close()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Fatal(err)
	}

	var sendingOrder structs.Order

	json.Unmarshal(byteData, &sendingOrder)
	cache.Set(sendingOrder)

	// Отправка реквеста
	req, err := http.NewRequest("GET", "/getOrder/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"orderId": ""})

	rr := httptest.NewRecorder()
	handler := GetOrder{Cache: cache}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response Response

	json.Unmarshal(rr.Body.Bytes(), &response)

	rightMessage := "ORDERS_NOT_FOUND"
	if response.Message != rightMessage {
		t.Errorf("handler returned unexpected message: got %v want %v",
			response.Message, rightMessage)
	}
}
