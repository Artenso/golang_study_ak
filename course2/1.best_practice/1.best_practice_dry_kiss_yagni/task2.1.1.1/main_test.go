package main

import (
	"reflect"
	"testing"
)

var product = &Product{
	ProductID:     ProductCocaCola,
	Sells:         []float64{10, 20, 30},
	Buys:          []float64{5, 15, 25},
	CurrentPrice:  35,
	ProfitPercent: 10,
}

func TestStatisticProfit_Average(t *testing.T) {
	type args struct {
		prices []float64
	}
	tests := []struct {
		name string
		s    *StatisticProfit
		args args
		want float64
	}{
		{name: "#1", args: args{prices: product.Sells}, want: float64(20)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Average(tt.args.prices); got != tt.want {
				t.Errorf("StatisticProfit.Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetAllData(t *testing.T) {
	tests := []struct {
		name string
		s    *StatisticProfit
		want []float64
	}{
		{
			name: "#1",
			s:    &StatisticProfit{product: product},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetAllData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatisticProfit.GetAllData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetAverageProfit(t *testing.T) {
	tests := []struct {
		name string
		s    *StatisticProfit
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetAverageProfit(); got != tt.want {
				t.Errorf("StatisticProfit.GetAverageProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetAverageProfitPercent(t *testing.T) {
	tests := []struct {
		name string
		s    *StatisticProfit
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetAverageProfitPercent(); got != tt.want {
				t.Errorf("StatisticProfit.GetAverageProfitPercent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetCurrentProfit(t *testing.T) {
	tests := []struct {
		name string
		s    *StatisticProfit
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCurrentProfit(); got != tt.want {
				t.Errorf("StatisticProfit.GetCurrentProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_GetDifferenceProfit(t *testing.T) {
	tests := []struct {
		name string
		s    *StatisticProfit
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetDifferenceProfit(); got != tt.want {
				t.Errorf("StatisticProfit.GetDifferenceProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatisticProfit_SetProduct(t *testing.T) {
	type args struct {
		p *Product
	}
	tests := []struct {
		name string
		s    *StatisticProfit
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetProduct(tt.args.p)
		})
	}
}

func TestStatisticProfit_Sum(t *testing.T) {
	type args struct {
		prices []float64
	}
	tests := []struct {
		name string
		s    *StatisticProfit
		args args
		want float64
	}{
		{name: "#1", args: args{prices: []float64{10, 20, 30}}, want: float64(60)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sum(tt.args.prices); got != tt.want {
				t.Errorf("StatisticProfit.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStatisticProfit(t *testing.T) {
	type args struct {
		opts []func(*StatisticProfit)
	}
	tests := []struct {
		name string
		args args
		want Profitable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatisticProfit(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatisticProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAverageProfit(t *testing.T) {
	type args struct {
		s *StatisticProfit
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithAverageProfit(tt.args.s)
		})
	}
}

func TestWithAverageProfitPercent(t *testing.T) {
	type args struct {
		s *StatisticProfit
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithAverageProfitPercent(tt.args.s)
		})
	}
}

func TestWithCurrentProfit(t *testing.T) {
	type args struct {
		s *StatisticProfit
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithCurrentProfit(tt.args.s)
		})
	}
}

func TestWithDifferenceProfit(t *testing.T) {
	type args struct {
		s *StatisticProfit
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithDifferenceProfit(tt.args.s)
		})
	}
}

func TestWithAllData(t *testing.T) {
	type args struct {
		s *StatisticProfit
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithAllData(tt.args.s)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
