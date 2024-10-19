package main

import (
	"fmt"
	"log"
	"time"

	client "github.com/Artenso/strategy/internal/client/exmo"
	"github.com/Artenso/strategy/internal/model"
)

type Exchanger interface {
	GetTicker() (model.Ticker, error)
	GetTrades(pairs ...string) (model.Trades, error)
	GetOrderBook(limit int, pairs ...string) (model.OrderBook, error)
	GetCurrencies() (model.Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (model.CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type GeneralIndicatorer interface {
	GetData(pair string, limit int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}
type Indicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type IndicatorEMA struct {
	client Exchanger
}

func NewIndicatorEMA(exmoClient Exchanger) Indicatorer {
	return &IndicatorEMA{
		client: exmoClient,
	}
}

func (i *IndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	data, err := i.client.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}

	res := calculateEMA(data, period)
	return res, nil
}

type IndicatorSMA struct {
	client Exchanger
}

func NewIndicatorSMA(exmoClient Exchanger) Indicatorer {
	return &IndicatorSMA{
		client: exmoClient,
	}
}

func (i *IndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	data, err := i.client.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}

	res := calculateSMA(data, period)
	return res, nil
}

// Функция для расчета простого скользящего среднего (SMA)
func calculateSMA(data []float64, period int) []float64 {
	var sma = make([]float64, len(data)/period)
	for i := range sma {
		sum := 0.0
		for _, d := range data[i*period : i*period+period] {
			sum += d
		}
		sma[i] = sum / float64(period)
	}
	return sma
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func calculateEMA(data []float64, period int) []float64 {
	if len(data) == 0 || period <= 0 {
		return nil
	}
	alpha := 2 / (float64(period) + 1.0)
	ema := make([]float64, len(data))
	ema[0] = data[0]
	for i := 1; i < len(data); i++ {
		ema[i] = alpha*data[i] + (1-alpha)*ema[i-1]
	}
	return ema
}

type GeneralIndicator struct{}

func (gi *GeneralIndicator) GetData(pair string, limit int, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, limit, 3, from, to)
}

func main() {
	// var exchange Exchanger
	exchange := client.NewExmo()

	indicatorSMA := NewIndicatorSMA(exchange)

	generalIndicator := &GeneralIndicator{}

	sma, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorSMA)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sma)

	indicatorEMA := NewIndicatorEMA(exchange)

	ema, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorEMA)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ema)
}
