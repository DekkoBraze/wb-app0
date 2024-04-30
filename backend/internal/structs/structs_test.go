package structs

import "testing"

func TestOrderValidatedCorrectOrder(t *testing.T) {
	order := Order{
		Order_uid: "order_uid",
		Items:     []Item{{Chrt_id: 5}},
		Payment:   Payment{Transaction: "transaction"},
		Delivery:  Delivery{Name: "name"},
	}

	result := OrderValidated(order)
	if !result {
		t.Errorf("handler returned unexpected message: got %v want %v",
			result, true)
	}
}

func TestOrderValidatedEmptyOrder(t *testing.T) {
	order := Order{}

	result := OrderValidated(order)
	if result {
		t.Errorf("handler returned unexpected message: got %v want %v",
			result, false)
	}
}

func TestOrderValidatedWithIdOnly(t *testing.T) {
	order := Order{Order_uid: "order_uid"}

	result := OrderValidated(order)
	if result {
		t.Errorf("handler returned unexpected message: got %v want %v",
			result, false)
	}
}

func TestOrderValidatedWithDeliveryOnly(t *testing.T) {
	order := Order{Delivery:  Delivery{Name: "name"}}

	result := OrderValidated(order)
	if result {
		t.Errorf("handler returned unexpected message: got %v want %v",
			result, false)
	}
}