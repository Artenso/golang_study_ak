package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mattevans/dinero"
)

func main() {
	// Пример использования функции
	rate, err := currencyPairRate("USD", "EUR", 100.0)
	if err != nil {
		log.Fatalf("some error: %s", err.Error())
	}
	fmt.Println(rate) // 92.3407
}

func currencyPairRate(from, to string, sum float64) (float64, error) {

	client := dinero.NewClient(
		"547171b3743449c58056938cb4cafe5d",
		from,
		20*time.Minute,
	)

	rsp, err := client.Rates.Get(to)
	if err != nil {
		return 0.0, err
	}

	return *rsp * sum, nil
}
