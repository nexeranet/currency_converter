package currency_converter

import "net/url"

type WTMineParams struct {
	HashRate         string
	Fees             string
	Power            string
	Cost             string
	HardwareCost     string
	BlockRewardValue *string
	DifficultyValue  *string
	BlockReward      *string
	Difficulty       *string
	ExchangeRate     *string
	BTCValue         *string
}

func (config WTMineParams) Query() string {
	//?hr=88.5&p=420.0&fee=1.0&cost=0.1&hcost=2.0&span_br=3&span_d=24
	query := url.Values{}
	//Hash rate
	query.Add("hr", config.HashRate)
	//Fees
	query.Add("fee", config.Fees)
	//Power
	query.Add("p", config.Power)
	// Cost
	query.Add("cost", config.Cost)
	//Hardware cost
	query.Add("hcost", config.HardwareCost)
	//Block reward value
	if config.BlockRewardValue != nil {
		query.Add("span_br", *config.BlockRewardValue)
	}
	// Difficulty value
	if config.DifficultyValue != nil {
		query.Add("span_d", *config.DifficultyValue)
	}

	//Block reward
	if config.BlockReward != nil {
		query.Add("br_enabled", "true")
		query.Add("br", *config.BlockReward)
	}

	//Difficulty
	if config.Difficulty != nil {
		query.Add("d_enabled", "true")
		query.Add("d", *config.Difficulty)
	}

	//Exchange rate
	//er_enabled=true&er=0.06778400&
	if config.ExchangeRate != nil {
		query.Add("er_enabled", "true")
		query.Add("er", *config.ExchangeRate)
	}

	//BTC value
	//btc_enabled=true&btc=36856.08
	if config.BTCValue != nil {
		query.Add("btc_enabled", "true")
		query.Add("btc", *config.BTCValue)
	}
	query.Add("commit", "Calculate")
	return query.Encode()
}
