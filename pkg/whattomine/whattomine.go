package whattomine

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	packageName string = "WHATTOMINEAPI"
)

var logInfo = log.New(os.Stdout, fmt.Sprintf("%s\t", packageName), log.Ldate|log.Ltime|log.Lshortfile)

type WhatToMineApi struct {
	Client     *http.Client
	Url        string
	Dictionary CalculatorsMap
	Ticker     *time.Ticker
	Coins      CoinsMap
}

func NewWhatToMineApi() *WhatToMineApi {
	return &WhatToMineApi{
		Url: "https://whattomine.com",
		Client: &http.Client{
			Timeout: 6 * time.Second,
		},
		Dictionary: make(CalculatorsMap),
		Coins:      make(CoinsMap),
	}
}

func (w *WhatToMineApi) Setup(wg *sync.WaitGroup) {
	defer wg.Done()
	err := w.UpdateDictionary()
	if err != nil {
		logInfo.Println("Error: ")
		return
	}
	w.UpdateCoins()
	w.CreateTicker()
}

func (w *WhatToMineApi) CreateTicker() {
	w.Ticker = time.NewTicker(1 * time.Minute)
	go func() {
		for range w.Ticker.C {
			err := w.UpdateDictionary()
			if err != nil {
				logInfo.Printf("Error in Ticker goroutine, %s", err.Error())
				return
			}
			w.UpdateCoins()
		}
	}()
}

func (w *WhatToMineApi) UpdateCoins() {
	logInfo.Println("Update coins started")
	for _, value := range w.Dictionary {
		time.Sleep(600 * time.Millisecond)
		coin, err := w.GetCoinById(value.Id)
		if err != nil {
			logInfo.Printf("Error can't get coin - %s:%d, %s", value.Tag, value.Id, err.Error())
			continue
		} else {
			w.Coins[coin.Tag] = coin
		}
	}
	logInfo.Println("Update coins finished")
}

func (w *WhatToMineApi) SetDictionary(dictionary CalculatorsMap) {
	w.Dictionary = dictionary
}

func (w *WhatToMineApi) UpdateDictionary() error {
	dictionary := make(CalculatorsMap)
	calc, err := w.GetCalculators()
	if err != nil {
		return err
	}
	for _, value := range calc.Coins {
		if value.Status != "Active" {
			continue
		}
		dictionary[value.Tag] = value
	}
	logInfo.Printf("Dictionary length - %d", len(dictionary))
	w.SetDictionary(dictionary)
	return nil
}

func (w *WhatToMineApi) FindCoinByTag(tag string) (Coin, error) {
	coin, ok := w.Coins[tag]
	if !ok {
		return Coin{}, &WError{
			Err: fmt.Errorf("Not found in dictionary - %s", tag),
		}
	}
	return coin, nil
}

func (w *WhatToMineApi) GetNetInfo(tag string) (Coin, error) {
	return w.FindCoinByTag(tag)
}
