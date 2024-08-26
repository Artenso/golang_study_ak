package client

import (
	"fmt"
	"time"
)

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	cHistory, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	candles := cHistory.Candles

	result := make([]float64, 0, len(candles))

	for _, candle := range candles {
		result = append(result, candle.C)
	}

	return result, nil
}
