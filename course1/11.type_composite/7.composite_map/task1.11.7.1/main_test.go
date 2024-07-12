package main

import (
	"reflect"
	"testing"
)

var testData = "Яблоки, яблоки, яблоки вишня черешня! Вишня в саду)"

func Test_countWordOccurrences(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{name: "#1", args: args{text: testData}, want: map[string]int{"яблоки": 3, "вишня": 2, "черешня": 1, "в": 1, "саду": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWordOccurrences(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWordOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
