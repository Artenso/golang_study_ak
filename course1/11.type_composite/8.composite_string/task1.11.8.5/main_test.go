package main

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{str: "Hello World!"}, want: "!dlroW olleH"},
		{name: "#2", args: args{str: "12345"}, want: "54321"},
		{name: "#3", args: args{str: ""}, want: ""},
		{name: "#4", args: args{str: "i"}, want: "i"},
		{name: "#5", args: args{str: "       "}, want: "       "},
		{name: "#6", args: args{str: ",! /.&?"}, want: "?&./ !,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.str); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
