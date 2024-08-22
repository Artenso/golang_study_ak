package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Artenso/oop-client/internal/model"
)

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (model.OrderBook, error) {
	url, err := url.JoinPath(e.url, orderBook)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	method := "POST"

	reqBody := strings.NewReader(
		fmt.Sprintf(
			"pair=%s&limit=%d",
			strings.Join(pairs, ","),
			limit,
		),
	)

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

	result := make(model.OrderBook)

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return result, nil
}
