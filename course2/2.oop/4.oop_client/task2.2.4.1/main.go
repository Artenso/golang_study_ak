package main

import (
	"fmt"
	"log"
	"time"

	client "github.com/Artenso/oop-client/internal/client/exmo"
)

func main() {
	exchange := client.NewExmo()
	ticker, err := exchange.GetCandlesHistory("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(ticker)
}
