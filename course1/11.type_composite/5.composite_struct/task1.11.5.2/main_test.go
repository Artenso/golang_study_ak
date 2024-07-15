package main

import (
	"fmt"
	"reflect"
	"testing"
)

var animals = []Animal{
	{Type: "Cat", Name: "Barsik", Age: 2},
	{Type: "Dog", Name: "Reks", Age: 3},
	{Type: "Horse", Name: "Spirit", Age: 1},
}

func Test_getAnimals(t *testing.T) {
	tests := []struct {
		name string
		want []Animal
	}{
		{
			name: "#1",
			want: animals},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAnimals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAnimals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preparePrint(t *testing.T) {
	type args struct {
		animals []Animal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{animals: animals},
			want: fmt.Sprintf(
				"Тип: %s, Имя: %s, Возраст: %d\nТип: %s, Имя: %s, Возраст: %d\nТип: %s, Имя: %s, Возраст: %d",
				animals[0].Type, animals[0].Name, animals[0].Age,
				animals[1].Type, animals[1].Name, animals[1].Age,
				animals[2].Type, animals[2].Name, animals[2].Age,
			)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preparePrint(tt.args.animals); got != tt.want {
				t.Errorf("preparePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
