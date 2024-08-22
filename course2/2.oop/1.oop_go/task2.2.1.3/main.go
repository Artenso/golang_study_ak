package main

import (
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

func (di *DineInOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("invalid quantity")
	}
	di.orderDetails[item] = quantity
	return nil
}

func (di *DineInOrder) RemoveItem(item string) error {
	if _, ok := di.orderDetails[item]; ok {
		delete(di.orderDetails, item)
		return nil
	}
	return fmt.Errorf("item not found")
}

func (di *DineInOrder) GetOrderDetails() map[string]int {
	return di.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (ta *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("invalid quantity")
	}
	ta.orderDetails[item] = quantity
	return nil
}

func (ta *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := ta.orderDetails[item]; ok {
		delete(ta.orderDetails, item)
		return nil
	}
	return fmt.Errorf("item not found")
}
func (ta *TakeAwayOrder) GetOrderDetails() map[string]int {
	return ta.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}
	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
