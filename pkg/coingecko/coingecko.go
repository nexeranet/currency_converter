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
	packageName        string = "COINGECKOAPI"
	CoinsInterval      int    = 1000
	DictionaryInterval int    = 5
)

var logInfo = log.New(os.Stdout, fmt.Sprintf("%s\t", packageName), log.Ldate|log.Ltime|log.Lshortfile)

type Coingecko struct {
	VsCurrencies    ConversionDictionary
	Coins           CoinsMap
	Client          *cg.Client
	TickerCoins     *time.Ticker
	TickerDictonary *time.Ticker
}

func NewCoingecko() *Coingecko {
	return &Coingecko{
		Client: cg.NewClient(&http.Client{
			Timeout: time.Second * 10,
		}),
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

func (c *Coingecko) CreateTickers() {
	c.TickerDictonary = time.NewTicker(time.Duration(DictionaryInterval) * time.Minute)
	go func() {
		for range c.TickerDictonary.C {
			ping, err := c.Ping()
			if err != nil {
				logInfo.Printf(err.Error())
				continue
			}
			logInfo.Printf("GeckoSays, %s", ping.GeckoSays)
			err = c.UpdateConvertiesDictionary()
			if err != nil {
				logInfo.Printf(err.Error())
				return
			}
			err = c.UpdateCoins()
			if err != nil {
				logInfo.Printf(err.Error())
				return
			}
		}
	}()
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
	logInfo.Printf("Coins dictonary is updated, len - %d \n", len(dictonary))
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
	logInfo.Printf("VsCurrencies dictonary is updated, len - %d \n", len(dictonary))
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
		for m_id, price := range map_prices {
			id_c := converstsMap[m_id]
			prices[id][id_c] = price
		}
	}
	return prices, nil
}
