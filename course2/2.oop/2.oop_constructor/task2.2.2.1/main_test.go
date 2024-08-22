package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		id   int
		opts []OrderOption
	}

	date := time.Now()

	tests := []struct {
		name string
		args args
		want *Order
	}{
		{
			name: "#1",
			args: args{
				id: 1,
				opts: []OrderOption{
					WithCustomerID("123"),
					WithItems([]string{"item1", "item2"}),
					WithOrderDate(date),
				},
			},
			want: &Order{
				ID:         1,
				CustomerID: "123",
				Items:      []string{"item1", "item2"},
				OrderDate:  date,
			},
		},
		{
			name: "#2",
			args: args{
				id: 1,
				opts: []OrderOption{
					WithCustomerID("123"),
				},
			},
			want: &Order{
				ID:         1,
				CustomerID: "123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(tt.args.id, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "#1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
