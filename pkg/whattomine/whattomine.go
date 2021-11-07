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
	packageName    string = "WHATTOMINEAPI"
	MinuteInterval int    = 5
	URL            string = "https://whattomine.com"
)

var logInfo = log.New(os.Stdout, fmt.Sprintf("%s\t", packageName), log.Ldate|log.Ltime|log.Lshortfile)

type WhatToMineApi struct {
	Client     *http.Client
	Url        string
	Dictionary CalculatorsMap
	Ticker     *time.Ticker
}

func NewWhatToMineApi() *WhatToMineApi {
	return &WhatToMineApi{
		Url: URL,
		Client: &http.Client{
			Timeout: 6 * time.Second,
		},
		Dictionary: make(CalculatorsMap),
	}
}

func (w *WhatToMineApi) Setup(wg *sync.WaitGroup) {
	defer wg.Done()
	err := w.UpdateDictionary()
	if err != nil {
		logInfo.Println("Error: ")
		return
	}
}

func (w *WhatToMineApi) CreateTickers() {
	w.Ticker = time.NewTicker(time.Duration(MinuteInterval) * time.Minute)
	go func() {
		for range w.Ticker.C {
			err := w.UpdateDictionary()
			if err != nil {
				logInfo.Printf("Error in Ticker goroutine, %s", err.Error())
				return
			}
		}
	}()
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

func (w *WhatToMineApi) GetCoinByTag(tag string) (Coin, error) {
	var result Coin
	calc, ok := w.Dictionary[tag]
	if !ok {
		return result, fmt.Errorf("Not found in whotomine dictionary - %s", tag)
	}
	coin, err := w.GetCoinById(calc.Id)
	if err != nil {
		return result, err
	}
	return coin, nil
}
