 <pre>
                                          _
 _ __   _____  _____ _ __ __ _ _ __   ___| |_
| '_ \ / _ \ \/ / _ \ '__/ _` | '_ \ / _ \ __|
| | | |  __/>  <  __/ | | (_| | | | |  __/ |_
|_| |_|\___/_/\_\___|_|  \__,_|_| |_|\___|\__|
</pre>
# Currency Converter API Client(Coingecko, Whattomine) for Go
## Resurses
Coingecko [link](https://www.coingecko.com/ru/api/documentation) - 100 requests/minute   
Whattomine [link](https://whattomine.com/coins)  - 80 requests/minute   
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
