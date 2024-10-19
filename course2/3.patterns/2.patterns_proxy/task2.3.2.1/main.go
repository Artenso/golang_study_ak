package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cinar/indicator"
)

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	MACD() ([]float64, []float64)
	EMA() []float64
}

func UnmarshalKLines(data []byte) (KLines, error) {
	var r KLines
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *KLines) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type KLines struct {
	Pair    string   `json:"pair"`
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *Lines) StochPrice() ([]float64, []float64) {
	k, d := indicator.StochasticOscillator(t.high, t.low, t.closing)

	return k, d
}

func (t *Lines) RSI(period int) ([]float64, []float64) {
	rs, rsi := indicator.RsiPeriod(period, t.closing)

	return rs, rsi
}

func (t *Lines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	_, rsi := t.RSI(rsiPeriod)
	k, d := indicator.StochasticOscillator(rsi, rsi, rsi)

	return k, d
}

func (t *Lines) MACD() ([]float64, []float64) {
	return indicator.Macd(t.closing)
}

func (t *Lines) EMA() []float64 {
	return indicator.Ema(5, t.closing)
}

type LinesProxy struct {
	lines Indicator
	cache map[string][]float64
}

func (l *LinesProxy) StochPrice() ([]float64, []float64) {
	k, kOk := l.cache["k_stochprice"]
	d, dOk := l.cache["d_stochprice"]
	if kOk && dOk {
		return k, d
	}

	k, d = l.lines.StochPrice()

	l.cache["k_stochprice"] = k
	l.cache["d_stochprice"] = d

	return k, d
}

func (l *LinesProxy) RSI(period int) ([]float64, []float64) {
	rs, rsOk := l.cache[fmt.Sprintf("rs_%v", period)]
	rsi, rsiOk := l.cache[fmt.Sprintf("rsi_%v", period)]
	if rsOk && rsiOk {
		return rs, rsi
	}

	rs, rsi = l.lines.RSI(period)

	l.cache[fmt.Sprintf("rs_%v", period)] = rs
	l.cache[fmt.Sprintf("rsi_%v", period)] = rsi

	return rs, rsi
}

func (l *LinesProxy) StochRSI(rsiPeriod int) ([]float64, []float64) {
	k, kOk := l.cache[fmt.Sprintf("k_stochrsi_%v", rsiPeriod)]
	d, dOk := l.cache[fmt.Sprintf("d_stochrsi_%v", rsiPeriod)]
	if kOk && dOk {
		return k, d
	}

	k, d = l.lines.StochPrice()

	l.cache[fmt.Sprintf("k_stochrsi_%v", rsiPeriod)] = k
	l.cache[fmt.Sprintf("d_stochrsi_%v", rsiPeriod)] = d

	return k, d
}

func (l *LinesProxy) MACD() ([]float64, []float64) {
	makd, makdOk := l.cache["makd"]
	signal, signalOk := l.cache["signal"]
	if makdOk && signalOk {
		return makd, signal
	}

	makd, signal = l.lines.MACD()

	l.cache["makd"] = makd
	l.cache["signal"] = signal

	return makd, signal
}

func (l *LinesProxy) EMA() []float64 {
	if ema, ok := l.cache["ema"]; ok {
		return ema
	}

	ema := l.lines.EMA()

	l.cache["ema"] = ema

	return ema
}

func LoadKlinesProxy(data []byte) *LinesProxy {
	return &LinesProxy{
		lines: LoadKlines(data),
		cache: make(map[string][]float64),
	}
}

func LoadKlines(data []byte) *Lines {
	klines, err := UnmarshalKLines(data)
	if err != nil {
		log.Fatal(err)
	}

	t := &Lines{}

	for _, v := range klines.Candles {
		t.closing = append(t.closing, v.C)
		t.low = append(t.low, v.L)
		t.high = append(t.high, v.H)
	}

	return t
}

func LoadCandles(pair string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.exmo.com/v1.1/candles_history?symbol=%s&resolution=30&from=1703056979&to=1705476839", pair), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	pair := "BTC_USD"
	candles := LoadCandles(pair)
	lines := LoadKlinesProxy(candles)
	fmt.Println(lines.RSI(3))
}
