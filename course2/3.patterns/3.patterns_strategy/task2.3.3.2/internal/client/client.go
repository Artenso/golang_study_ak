package client

import (
	"time"

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
