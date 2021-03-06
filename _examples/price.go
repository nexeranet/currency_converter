package main

import (
	"fmt"
	"time"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	price, err := converter.GetPrice("BTC", "USD")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(price)
	prices, err := converter.GetPricesInGroups([]string{"BTC"}, []string{"ETH", "USD"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(prices)
	time.Sleep(10 * time.Minute)
	fmt.Println("END")
}
