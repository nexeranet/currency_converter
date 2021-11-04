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
	Name             string      `json:"name,omitempty"`
	Algorithm        string      `json:"algorithm"`
	BlockTime        json.Number `json:"block_time"`
	BlockReward      float32     `json:"block_reward"`
	BlockReward24    float32     `json:"block_reward24"`
	BlockReward3     float32     `json:"block_reward3,omitempty"`
	BlockReward7     float32     `json:"block_reward7,omitempty"`
	LastBlock        float32     `json:"last_block"`
	Difficulty       float32     `json:"difficulty"`
	Difficulty24     float32     `json:"difficulty24"`
	Difficulty3      float32     `json:"difficulty3,omitempty"`
	Difficulty7      float32     `json:"difficulty7,omitempty"`
	Nethash          float64     `json:"nethash"`
	ExchangeRate     float32     `json:"exchange_rate"`
	ExchangeRate24   float32     `json:"exchange_rate24"`
	ExchangeRate3    float32     `json:"exchange_rate3,omitempty"`
	ExchangeRate7    float32     `json:"exchange_rate7,omitempty"`
	ExchangeRateVol  float32     `json:"exchange_rate_vol"`
	ExchangeRateCurr string      `json:"exchange_rate_curr"`
	MarketCap        string      `json:"market_cap"`
	PoolFee          string      `json:"pool_fee,omitempty"`
	EstimatedRewards string      `json:"estimated_rewards,omitempty"`
	BtcRevenue       string      `json:"btc_revenue,omitempty"`
	Revenue          string      `json:"revenue,omitempty"`
	Cost             string      `json:"cost,omitempty"`
	Profit           string      `json:"profit,omitempty"`
	Status           string      `json:"status,omitempty"`
	Lagging          bool        `json:"lagging,omitempty"`
	Testing          bool        `json:"testing,omitempty"`
	Listed           bool        `json:"listed,omitempty"`
	Timestamp        int32       `json:"timestamp,omitempty"`

	EstimatedRewards24 string `json:"estimated_rewards24,omitempty"`
	BtcRevenue24       string `json:"btc_revenue24,omitempty"`
	Profitability      int    `json:"profitability,omitempty"`
	Profitability24    int    `json:"profitability24,omitempty"`
}

type WError struct {
	Err error
}

func (w *WError) Error() string {
	return fmt.Sprintf("%s : %s", packageName, w.Err)
}
