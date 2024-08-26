package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Artenso/strategy/internal/model"
)

func (e *Exmo) GetTrades(pairs ...string) (model.Trades, error) {
	url, err := url.JoinPath(e.url, trades)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	method := "POST"

	result := make(model.Trades)

	for _, pair := range pairs {
		reqBody := strings.NewReader(pair)

		req, err := http.NewRequest(method, url, reqBody)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := e.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}

		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	}

	return result, nil
}
