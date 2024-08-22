package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(customerID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = customerID
	}
}

func WithItems(items []string) OrderOption {
	return func(o *Order) {
		o.Items = items
	}
}

func WithOrderDate(date time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = date
	}
}

func NewOrder(id int, opts ...OrderOption) *Order {
	order := &Order{
		ID: id,
	}

	for _, opt := range opts {
		opt(order)
	}

	return order
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))
	fmt.Printf("Order: %+v\n", order)
}
