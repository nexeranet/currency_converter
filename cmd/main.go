package main

import (
	"fmt"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	err := converter.Setup()
	if err != nil {
		fmt.Println(err.Error())
	}
	coin, err := converter.GetNetInfo("BTC")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(coin)
	fmt.Println("END")
}
