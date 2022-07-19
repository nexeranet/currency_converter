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

func NewConverter(flags ...bool) *Converter {
	debug := false
	if len(flags) > 0 {
		debug = flags[0]
	}
	return &Converter{
		WhattomineApi: whattomine.NewWhatToMineApi(debug),
		CoingeckoApi:  coingecko.NewCoingecko(debug),
	}
}

func (c *Converter) Setup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go c.CoingeckoApi.Setup(&wg)
	go c.WhattomineApi.Setup(&wg)
	wg.Wait()
	fmt.Println("###################\t Setup is done \t##########################")
}

func (c *Converter) CreateTickers() {
	c.WhattomineApi.CreateTickers()
	c.CoingeckoApi.CreateTickers()
}
func (c *Converter) StopTickers() {
	c.WhattomineApi.Ticker.Stop()
	c.CoingeckoApi.Ticker.Stop()
}

// WHATTOMINEAPI
func (c *Converter) GetNetInfo(tag string) (whattomine.Coin, error) {
	return c.WhattomineApi.GetCoinByTag(tag)
}

func (c *Converter) GetWTCoin(tag string) (whattomine.Coin, error) {
	return c.WhattomineApi.GetCoinByTag(tag)
}

func (c *Converter) GetNetInfoWithParams(tag string, config WTMineParams) (whattomine.Coin, error) {
	return c.WhattomineApi.GetCoinByTagWithCalculate(tag, config.Query())
}

func (c *Converter) GetNetInfoWithQuery(tag string, query string) (whattomine.Coin, error) {
	return c.WhattomineApi.GetCoinByTagWithCalculate(tag, query)
}

// COINGECKOAPI

func (c *Converter) GetPrice(name, convert_name string) (float32, error) {
	return c.CoingeckoApi.GetPrice(name, convert_name)
}

func (c *Converter) GetPricesInGroups(names, converts []string) (coingecko.PricesMap, error) {
	return c.CoingeckoApi.GetPricesInGroups(names, converts)
}
