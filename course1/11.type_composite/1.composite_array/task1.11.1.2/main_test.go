package main

import (
	"reflect"
	"testing"
)

var intArr = [8]int{5, 2, 8, 1, 9, 3, 7, 4}
var floatArr = [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}

func TestSortDescInt(t *testing.T) {
	tests := []struct {
		name string
		in   [8]int
		want [8]int
	}{
		{name: "1", in: intArr, want: [8]int{9, 8, 7, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescInt(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortAscInt(t *testing.T) {
	tests := []struct {
		name string
		in   [8]int
		want [8]int
	}{
		{name: "1", in: intArr, want: [8]int{1, 2, 3, 4, 5, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAscInt(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAscInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortDescFloat(t *testing.T) {
	tests := []struct {
		name string
		in   [8]float64
		want [8]float64
	}{
		{name: "1", in: floatArr, want: [8]float64{9.9, 8.8, 7.7, 5.5, 4.4, 3.3, 2.2, 1.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescFloat(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortAscFloat(t *testing.T) {
	tests := []struct {
		name string
		in   [8]float64
		want [8]float64
	}{
		{name: "1", in: floatArr, want: [8]float64{1.1, 2.2, 3.3, 4.4, 5.5, 7.7, 8.8, 9.9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortAscFloat(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAscFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
