package main

import (
	"fmt"
	"time"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	coin, err := converter.GetNetInfoWithParams("BTC", currency_converter.WTMineParams{})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(coin)
	time.Sleep(1 * time.Minute)
	fmt.Println("END")
}
