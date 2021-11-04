package main

import (
	"fmt"
	"time"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	coin, err := converter.GetNetInfo("BTC")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(coin)
	res, err := converter.WhattomineApi.GetCoinById(123)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
	time.Sleep(1 * time.Minute)
	fmt.Println("END")
}
