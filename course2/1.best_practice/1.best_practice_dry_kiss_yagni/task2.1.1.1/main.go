package main

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(prices []float64) float64
	Sum(prices []float64) float64
}

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

// Average implements Profitable.
func (s *StatisticProfit) Average(prices []float64) float64 {
	return s.Sum(prices) / float64(len(prices))
}

// GetAllData implements Profitable.
func (s *StatisticProfit) GetAllData() []float64 {
	return s.getAllData()
}

// GetAverageProfit implements Profitable.
func (s *StatisticProfit) GetAverageProfit() float64 {
	return s.getAverageProfit()
}

// GetAverageProfitPercent implements Profitable.
func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	return s.getAverageProfitPercent()
}

// GetCurrentProfit implements Profitable.
func (s *StatisticProfit) GetCurrentProfit() float64 {
	return s.getCurrentProfit()
}

// GetDifferenceProfit implements Profitable.
func (s *StatisticProfit) GetDifferenceProfit() float64 {
	return s.getDifferenceProfit()
}

// SetProduct implements Profitable.
func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

// Sum implements Profitable.
func (s *StatisticProfit) Sum(prices []float64) float64 {
	var summ float64

	for _, price := range prices {
		summ += price
	}

	return summ
}

func NewStatisticProfit(opts ...func(*StatisticProfit)) Profitable {
	sp := &StatisticProfit{}
	for _, opt := range opts {
		opt(sp)
	}
	return sp
}

// WithAverageProfit sets the average profit of the product
//
// Average Profit = Average Sells - Average Buys
func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
}

// WithAverageProfitPercent sets the average profit percent of the product
//
// Average Profit Percent = Average Profit / Average Buys * 100
func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		return s.getAverageProfit() / s.Average(s.product.Buys) * 100
	}
}

// WithCurrentProfit sets the current profit of the product
//
// Current Profit = Current Price - Current Price * (100 - Profit Percent) / 100
func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - s.product.CurrentPrice*(100-s.product.ProfitPercent)/100
	}
}

// WithDifferenceProfit sets the difference profit of the product
//
// Difference Profit = Current Price - Average Sells
func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		return s.product.CurrentPrice - s.Average(s.product.Sells)
	}
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfit != nil {
			res = append(res, s.getAverageProfit())
		}
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite
	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)
	statProfit.SetProduct(product)
}
