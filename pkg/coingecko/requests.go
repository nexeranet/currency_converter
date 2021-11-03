package coingecko

import (
	"github.com/superoo7/go-gecko/v3/types"
)

func (c *Coingecko) CoinsList() (*types.CoinList, error) {
	return c.Client.CoinsList()
}

func (c *Coingecko) SimpleSupportedVSCurrencies() (*types.SimpleSupportedVSCurrencies, error) {
	return c.Client.SimpleSupportedVSCurrencies()
}

func (c *Coingecko) SimpleSinglePrice(name, currency string) (*types.SimpleSinglePrice, error) {
	return c.Client.SimpleSinglePrice(name, currency)
}
func (c *Coingecko) SimplePrice(names, currencies []string) (*map[string]map[string]float32, error) {
	return c.Client.SimplePrice(names, currencies)
}
