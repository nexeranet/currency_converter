package main

import (
	"fmt"
	"time"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	price, err := converter.GetPrice("BTC", "ETHasdfsa")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(price)
	prices, err := converter.GetPricesInGroups([]string{"BTC"}, []string{"ETH", "USD"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(prices)
	coins, err := converter.WhattomineApi.GetCoins()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(coins)
	time.Sleep(10 * time.Minute)
	fmt.Println("END")
}
