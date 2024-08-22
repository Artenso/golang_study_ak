package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Artenso/oop-client/internal/model"
)

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (model.CandlesHistory, error) {

	url, err := url.JoinPath(e.url, candlesHistory)
	if err != nil {
		return model.CandlesHistory{}, fmt.Errorf(err.Error())
	}

	urlWithQParams := fmt.Sprintf(
		`%s?symbol=%s&resolution=%d&from=%d&to=%d`,
		url,
		pair,
		limit,
		start.Unix(),
		end.Unix(),
	)

	method := "GET"

	req, err := http.NewRequest(method, urlWithQParams, nil)
	if err != nil {
		return model.CandlesHistory{}, fmt.Errorf(err.Error())
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return model.CandlesHistory{}, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.CandlesHistory{}, fmt.Errorf(err.Error())
	}

	result := model.CandlesHistory{}

	if err := json.Unmarshal(body, &result); err != nil {
		return model.CandlesHistory{}, fmt.Errorf(err.Error())
	}

	return result, nil
}
