package coingecko

import (
	"fmt"
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
