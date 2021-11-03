package coingecko

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	cg "github.com/superoo7/go-gecko/v3"
)

const (
	packageName string = "COINGECKOAPI"
)

var logInfo = log.New(os.Stdout, fmt.Sprintf("%s\t", packageName), log.Ldate|log.Ltime|log.Lshortfile)

type Coingecko struct {
	VsCurrencies ConversionDictionary
	Coins        CoinsMap
	Client       *cg.Client
	Ticker       *time.Ticker
}

func NewCoingecko() *Coingecko {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return &Coingecko{
		Client:       cg.NewClient(httpClient),
		VsCurrencies: make(ConversionDictionary),
		Coins:        make(CoinsMap),
	}
}

func (c *Coingecko) Setup(wg *sync.WaitGroup) {
	defer wg.Done()
	err := c.UpdateConvertiesDictionary()
	if err != nil {
		logInfo.Println("Error: ", err.Error())
		return
	}
	err = c.UpdateCoins()
	if err != nil {
		logInfo.Println("Error ", err.Error())
		return
	}
}

func (c *Coingecko) UpdateCoins() error {
	coins, err := c.CoinsList()
	if err != nil {
		return &GKError{
			Err: err,
		}
	}
	dictonary := make(CoinsMap)
	for _, value := range *coins {
		dictonary[strings.ToUpper(value.Symbol)] = Coin{
			ID:     value.ID,
			Symbol: value.Symbol,
			Name:   value.Name,
		}
	}
	c.Coins = dictonary
	logInfo.Println("Coins dictonary is updated")
	return nil
}
func (c *Coingecko) UpdateConvertiesDictionary() error {
	converties, err := c.SimpleSupportedVSCurrencies()
	dictonary := make(ConversionDictionary)
	if err != nil {
		return &GKError{
			Err: err,
		}
	}
	for _, value := range *converties {
		dictonary[strings.ToUpper(value)] = value
	}
	c.VsCurrencies = dictonary
	logInfo.Println("VsCurrencies dictonary is updated")
	return nil
}

func (c *Coingecko) GetPrice(name, convert string) (float32, error) {
	name_id, ok := c.Coins[name]
	if !ok {
		return 0.0, &GKError{
			Err: fmt.Errorf("Not found in Coins dictionary (Coingecko doesn't have this currency) - %s", name),
		}
	}
	convert_id, ok := c.VsCurrencies[convert]
	if !ok {
		return 0.0, &GKError{
			Err: fmt.Errorf("Not found in VsCurrencies dictionary (Coingecko doesn't have this currency or can't convert) - %s", convert),
		}
	}
	coin, err := c.SimpleSinglePrice(name_id.ID, convert_id)
	if err != nil {
		return 0.0, &GKError{
			Err: err,
		}
	}
	return coin.MarketPrice, nil
}

func (c *Coingecko) GetPricesInGroups(names, conversts []string) (PricesMap, error) {
	var prices PricesMap
	namesMap := make(map[string]string)
	converstsMap := make(map[string]string)
	converts_ids := []string{}
	names_ids := []string{}
	for _, value := range names {
		id, ok := c.Coins[value]
		if !ok {
			return prices, &GKError{
				Err: fmt.Errorf("Not found in Coins dictionary (Coingecko doesn't have this currency) - %s", value),
			}
		}
		namesMap[id.ID] = value
		names_ids = append(names_ids, id.ID)
	}
	for _, value := range conversts {
		id, ok := c.VsCurrencies[value]
		if !ok {
			return prices, &GKError{
				Err: fmt.Errorf("Not found in VsCurrencies dictionary (Coingecko doesn't have this currency or can't convert) - %s", value),
			}
		}
		converts_ids = append(converts_ids, id)
		converstsMap[id] = value
	}
	data, err := c.SimplePrice(names_ids, converts_ids)
	if err != nil {
		return prices, &GKError{
			Err: err,
		}
	}
	prices = make(PricesMap)
	for key, map_prices := range *data {
		id := namesMap[key]
		prices[id] = make(map[string]float32)
		for m_id, price := map_prices {

		}
	}
	return prices, nil
}
