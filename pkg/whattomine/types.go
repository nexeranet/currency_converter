package whattomine

import (
	"encoding/json"
	"fmt"
)

type CoinsMap map[string]Coin

type Coins struct {
	Coins CoinsMap `json:"coins"`
}

type Calculator struct {
	Id        int    `json:"id"`
	Tag       string `json:"tag"`
	Algorithm string `json:"algorithm"`
	Name      string `json:"name"`
	Listed    bool   `json:"listed"`
	Status    string `json:"status"`
	Testing   bool   `json:"testing"`
	Lagging   bool   `json:"lagging"`
}

type CalculatorsMap map[string]Calculator

type Calculators struct {
	Coins CalculatorsMap `json:"coins"`
}

type Coin struct {
	Id               int         `json:"id"`
	Tag              string      `json:"tag"`
	Name             string      `json:"name"`
	Algorithm        string      `json:"algorithm"`
	BlockTime        json.Number `json:"block_time"`
	BlockReward      float32     `json:"block_reward"`
	BlockReward24    float32     `json:"block_reward24"`
	BlockReward3     float32     `json:"block_reward3"`
	BlockReward7     float32     `json:"block_reward7"`
	LastBlock        float32     `json:"last_block"`
	Difficulty       float32     `json:"difficulty"`
	Difficulty24     float32     `json:"difficulty24"`
	Difficulty3      float32     `json:"difficulty3"`
	Difficulty7      float32     `json:"difficulty7"`
	Nethash          float64     `json:"nethash"`
	ExchangeRate     float32     `json:"exchange_rate"`
	ExchangeRate24   float32     `json:"exchange_rate24"`
	ExchangeRate3    float32     `json:"exchange_rate3"`
	ExchangeRate7    float32     `json:"exchange_rate7"`
	ExchangeRateVol  float32     `json:"exchange_rate_vol"`
	ExchangeRateCurr string      `json:"exchange_rate_curr"`
	MarketCap        string      `json:"market_cap"`
	PoolFee          string      `json:"pool_fee"`
	EstimatedRewards string      `json:"estimated_rewards"`
	BtcRevenue       string      `json:"btc_revenue,omitempty"`
	Revenue          string      `json:"revenue"`
	Cost             string      `json:"cost"`
	Profit           string      `json:"profit"`
	Status           string      `json:"status"`
	Lagging          bool        `json:"lagging"`
	Testing          bool        `json:"testing"`
	Listed           bool        `json:"listed"`
	Timestamp        int32       `json:"timestamp"`
}

type WError struct {
	Err error
}

func (w *WError) Error() string {
	return fmt.Sprintf("%s : %s", packageName, w.Err)
}
