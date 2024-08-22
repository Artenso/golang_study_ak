package main

import "testing"

func Test_createUniqueText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{text: "bar bar bar foo foo baz"}, want: "bar foo baz"},
		{name: "#2", args: args{text: "bar foo bar foo baz bar"}, want: "bar foo baz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createUniqueText(tt.args.text); got != tt.want {
				t.Errorf("createUniqueText() = %v, want %v", got, tt.want)
			}
		})
	}
}
