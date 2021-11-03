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

func (c *Converter) GetNetInfo(tag string) (whattomine.Coin, error) {
	return c.WhattomineApi.GetNetInfo(tag)
}

func (c *Converter) GetPrice(name, convert_name string) (float32, error) {
	return c.CoingeckoApi.GetPrice(name, convert_name)
}

func (c *Converter) GetPricesInGroups(names, converts []string) (coingecko.PricesMap, error) {
	return c.CoingeckoApi.GetPricesInGroups(names, converts)
}
