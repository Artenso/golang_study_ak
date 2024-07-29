package main

import "testing"

func Test_generateMathString(t *testing.T) {
	type args struct {
		operands []int
		operator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{operands: []int{6, 2, 3}, operator: "+"}, want: "6 + 2 + 3 = 11"},
		{name: "#2", args: args{operands: []int{6, 2, 3}, operator: "-"}, want: "6 - 2 - 3 = 1"},
		{name: "#3", args: args{operands: []int{6, 2, 3}, operator: "/"}, want: "6 / 2 / 3 = 1.00"},
		{name: "#4", args: args{operands: []int{6, 2, 3}, operator: "*"}, want: "6 * 2 * 3 = 36"},
		{name: "#5", args: args{operands: []int{6, 2, 3}, operator: "%"}, want: "unknown operator"},
		{name: "#5", args: args{operands: []int{}, operator: "+"}, want: "not enough operands"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMathString(tt.args.operands, tt.args.operator); got != tt.want {
				t.Errorf("generateMathString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_summ(t *testing.T) {
	type args struct {
		operands []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{operands: []int{6, 2, 3}}, want: "6 + 2 + 3 = 11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summ(tt.args.operands); got != tt.want {
				t.Errorf("summ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	type args struct {
		operands []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{operands: []int{6, 2, 3}}, want: "6 - 2 - 3 = 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.operands); got != tt.want {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mult(t *testing.T) {
	type args struct {
		operands []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{operands: []int{6, 2, 3}}, want: "6 * 2 * 3 = 36"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mult(tt.args.operands); got != tt.want {
				t.Errorf("mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div(t *testing.T) {
	type args struct {
		operands []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{operands: []int{6, 2, 3}}, want: "6 / 2 / 3 = 1.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.args.operands); got != tt.want {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}
