package whattomine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	res, err := w.Client.Get(fmt.Sprintf("%s/%s", w.Url, path))
	if err != nil {
		return result, err
	}
	if res.StatusCode != http.StatusOK {
		return result, &WError{
			Err: fmt.Errorf("Status code : %d", res.StatusCode),
		}
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