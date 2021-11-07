package main

import (
	"fmt"
	"time"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	converter.CreateTickers()
	coin, err := converter.GetNetInfo("BTC")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(coin)
	time.Sleep(1 * time.Minute)
	converter.StopTickers()
	time.Sleep(5 * time.Minute)
	fmt.Println("END")
}
