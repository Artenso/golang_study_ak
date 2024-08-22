package main

import (
	"reflect"
	"testing"
)

var dish1 = Dish{Name: "Pizza", Price: 10.99}
var dish2 = Dish{Name: "Burger", Price: 5.99}

func TestOrder_AddDish(t *testing.T) {
	var order = Order{}
	tests := []struct {
		name  string
		order *Order
		want  []Dish
	}{
		{name: "#1", order: &order, want: []Dish{dish1, dish2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.order.AddDish(dish1)
			tt.order.AddDish(dish2)
			if !reflect.DeepEqual(order.Dishes, tt.want) {
				t.Errorf("order.Dishes = %v, want %v", order.Dishes, tt.want)
			}
		})
	}
}

func TestOrder_RemoveDish(t *testing.T) {
	var order = Order{
		Dishes: []Dish{dish1, dish2},
	}
	tests := []struct {
		name  string
		order *Order
		want  []Dish
	}{
		{name: "#1", order: &order, want: []Dish{dish2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.order.RemoveDish(dish1)
			if !reflect.DeepEqual(order.Dishes, tt.want) {
				t.Errorf("order.Dishes = %v, want %v", order.Dishes, tt.want)
			}
		})
	}
}

func TestOrder_CalculateTotal(t *testing.T) {
	var order = Order{
		Dishes: []Dish{dish1, dish2},
	}
	tests := []struct {
		name  string
		order *Order
		want  float64
	}{
		{name: "#1", order: &order, want: 16.98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.order.CalculateTotal()
			if order.Total != tt.want {
				t.Errorf("order.Total = %v, want %v", order.Total, tt.want)
			}
		})
	}
}
