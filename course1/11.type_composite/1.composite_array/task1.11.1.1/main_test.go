package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		in   [8]int
		want int
	}{
		{name: "1", in: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, want: 36},
		{name: "2", in: [8]int{2, 3, 4, 5, 6, 7, 8, 9}, want: 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.in); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_average(t *testing.T) {
	tests := []struct {
		name string
		in   [8]int
		want float64
	}{
		{name: "1", in: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, want: 4.5},
		{name: "2", in: [8]int{2, 3, 4, 5, 6, 7, 8, 9}, want: 5.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.in); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageFloat(t *testing.T) {
	tests := []struct {
		name string
		in   [8]float64
		want float64
	}{
		{name: "1", in: [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageFloat(tt.in); got != tt.want {
				t.Errorf("averageFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		in   [8]int
		want [8]int
	}{
		{name: "1", in: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, want: [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
		{name: "2", in: [8]int{2, 3, 4, 5, 6, 7, 8, 9}, want: [8]int{9, 8, 7, 6, 5, 4, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
