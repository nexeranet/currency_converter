package coingecko

import (
	"fmt"
	"sync"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
type ConversionDictionary map[string]string

type CoinsMap map[string]Coin

type PricesMap map[string]map[string]float32

type GKError struct {
	Err error
}

func (g *GKError) Error() string {
	return fmt.Sprintf("%s : %s", packageName, g.Err)
}

type DictionaryCoins struct {
	Mutex sync.Mutex
	Data  CoinsMap
}

func (d *DictionaryCoins) Swap(dict CoinsMap) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	d.Data = dict
}

func (d *DictionaryCoins) Get(tag string) (Coin, bool) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	record, ok := d.Data[tag]
	return record, ok
}

func (d *DictionaryCoins) Set(tag string, record Coin) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	d.Data[tag] = record
}

type DictionaryConversion struct {
	Mutex sync.Mutex
	Data  ConversionDictionary
}

func (d *DictionaryConversion) Swap(dict ConversionDictionary) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	d.Data = dict
}

func (d *DictionaryConversion) Get(tag string) (string, bool) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	record, ok := d.Data[tag]
	return record, ok
}

func (d *DictionaryConversion) Set(tag, record string) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	d.Data[tag] = record
}
