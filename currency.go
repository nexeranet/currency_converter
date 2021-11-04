/*

                                          _
 _ __   _____  _____ _ __ __ _ _ __   ___| |_
| '_ \ / _ \ \/ / _ \ '__/ _` | '_ \ / _ \ __|
| | | |  __/>  <  __/ | | (_| | | | |  __/ |_
|_| |_|\___/_/\_\___|_|  \__,_|_| |_|\___|\__|
*/
package currency_converter

import (
	"fmt"
	"sync"

	"github.com/nexeranet/currency_converter/pkg/coingecko"
	"github.com/nexeranet/currency_converter/pkg/whattomine"
)

type Converter struct {
	WhattomineApi *whattomine.WhatToMineApi
	CoingeckoApi  *coingecko.Coingecko
}

func NewConverter() *Converter {
	return &Converter{
		WhattomineApi: whattomine.NewWhatToMineApi(),
		CoingeckoApi:  coingecko.NewCoingecko(),
	}
}

func (c *Converter) Setup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go c.CoingeckoApi.Setup(&wg)
	go c.WhattomineApi.Setup(&wg)
	wg.Wait()
	c.WhattomineApi.CreateTickers()
	c.CoingeckoApi.CreateTickers()
	fmt.Println("###################\t Setup is done \t##########################")
}

// WHATTOMINEAPI
func (c *Converter) GetNetInfo(tag string) (whattomine.Coin, error) {
	return c.WhattomineApi.GetNetInfo(tag)
}

// func (c *Converter) GetWhattomineCalculators() (whattomine.Calculators, error) {
// return c.WhattomineApi.GetCalculators()
// }
//
// func (c *Converter) GetWhattomineCoins() (whattomine.Coins, error) {
// return c.WhattomineApi.GetCoins()
// }
//
// func (c *Converter) GetWhattomineCoinById(id int) (whattomine.Coin, error) {
// return c.WhattomineApi.GetCoinById(id)
// }

// COINGECKOAPI

func (c *Converter) GetPrice(name, convert_name string) (float32, error) {
	return c.CoingeckoApi.GetPrice(name, convert_name)
}

func (c *Converter) GetPricesInGroups(names, converts []string) (coingecko.PricesMap, error) {
	return c.CoingeckoApi.GetPricesInGroups(names, converts)
}
