package main

import (
	"reflect"
	"testing"
)

func Test_countRussianLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[rune]int
	}{
		{
			name: "#1",
			args: args{s: "Привет, Мир!"},
			want: map[rune]int{
				'п': 1,
				'р': 2,
				'и': 2,
				'в': 1,
				'е': 1,
				'т': 1,
				'м': 1,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRussianLetters(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countRussianLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRussianLetter(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#1", args: args{char: 'а'}, want: true},
		{name: "#2", args: args{char: 'я'}, want: true},
		{name: "#3", args: args{char: 'о'}, want: true},
		{name: "#4", args: args{char: 'g'}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRussianLetter(tt.args.char); got != tt.want {
				t.Errorf("isRussianLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
