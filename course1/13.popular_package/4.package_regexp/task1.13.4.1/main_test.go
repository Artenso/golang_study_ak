package main

import "testing"

func Test_isValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#1", args: args{email: "user@example.com"}, want: true},
		{name: "#2", args: args{email: "some.user@domain.co.uk"}, want: true},
		{name: "#3", args: args{email: "invalid_email"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEmail(tt.args.email); got != tt.want {
				t.Errorf("isValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
