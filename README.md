# Carrcency Converter API Client(Coingecko, Whattomine) for Go
## Resurses
Coingecko [link](https://www.coingecko.com/ru/api/documentation)  
Whattomine [link](https://whattomine.com/coins)  
## Usage
For usage, checkout [Example folder](https://github.com/nexeranet/currency_converter/tree/main/examples)
```go
package main

import (
	"fmt"
	"time"
	"github.com/nexeranet/currency_converter"
)

func main() {
	converter := currency_converter.NewConverter()
	converter.Setup()
	price, err := converter.GetPrice("BTC", "ETH")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(price)
	prices, err := converter.GetPricesInGroups([]string{"BTC"}, []string{"ETH", "USD"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(prices)
}
```
