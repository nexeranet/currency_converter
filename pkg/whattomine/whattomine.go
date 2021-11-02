package whattomine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WhatToMineApi struct {
	Url        string
	Dictionary map[string]int
}

func NewWhatToMineApi() *WhatToMineApi {
	return &WhatToMineApi{
		Url: "https://whattomine.com",
	}
}

func (w *WhatToMineApi) Setup() error {
	err := w.UpdateDictionary()
	if err != nil {
		return err
	}
	return nil
}

func (w *WhatToMineApi) SetDictionary(dictionary map[string]int) {
	w.Dictionary = dictionary
}
func (w *WhatToMineApi) UpdateDictionary() error {
	dictionary := make(map[string]int)
	calc, err := w.GetCalculators()
	if err != nil {
		return err
	}
	for _, value := range calc.Coins {
		dictionary[value.Tag] = value.Id
	}
	w.SetDictionary(dictionary)
	return nil
}

func (w *WhatToMineApi) GetCoinById(id int) (Coin, error) {
	var con Coin
	res, err := w.Get(fmt.Sprintf("coins/%d.json", id))
	if err != nil {
		return con, err
	}
	err = json.Unmarshal(res, &con)
	if err != nil {
		return con, err
	}
	return con, nil
}

func (w *WhatToMineApi) GetCoins() (Coins, error) {
	var coins Coins
	res, err := w.Get("coins.json")
	if err != nil {
		return coins, err
	}
	err = json.Unmarshal(res, &coins)
	if err != nil {
		return coins, err
	}
	return coins, nil
}

func (w *WhatToMineApi) Get(path string) ([]byte, error) {
	var result []byte
	res, err := http.Get(fmt.Sprintf("%s/%s", w.Url, path))
	if err != nil {
		return result, err
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return result, err
	}
	return content, nil
}

func (w *WhatToMineApi) GetCalculators() (Calculators, error) {
	var calc Calculators
	res, err := w.Get("calculators.json")
	if err != nil {
		return calc, err
	}
	err = json.Unmarshal(res, &calc)
	if err != nil {
		return calc, err
	}
	return calc, nil
}

func (w *WhatToMineApi) GetNetInfo(tag string) (Coin, error) {
	var result Coin
	id, ok := w.Dictionary[tag]
	if !ok {
		return result, fmt.Errorf("Not found in whotomine dictionary - %s", tag)
	}
	coin, err := w.GetCoinById(id)
	if err != nil {
		return result, err
	}
	return coin, nil
}
