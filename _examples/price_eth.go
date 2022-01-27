package main

import (
	"fmt"
	"time"

	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	price, err := converter.GetPrice("ETH", "USD")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(price)
	time.Sleep(10 * time.Minute)
	fmt.Println("END")
}
