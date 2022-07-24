package whattomine

import (
	"net/http"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/h2non/gock.v1"
)

var API = NewWhatToMineApi(false)

func Cleanup(t *testing.T) {
	t.Cleanup(func() {
		gock.EnableNetworking()
		gock.OffAll()
	})
	gock.DisableNetworking()
}

func TestGetCoinById(t *testing.T) {
	Cleanup(t)

	type input struct {
		id int
	}

	type output struct {
		res     Coin
		withErr bool
	}

	type test struct {
		name   string
		setup  func()
		input  input
		output output
	}

	tests := []test{
		{
			name: "200",
			setup: func() {
				gock.New(URL).
					Get("coins/1.json").
					Reply(http.StatusOK).
					File(path.Join("json", "coin.json"))
			},
			input: input{
				id: 1,
			},
			output: output{
				res: Coin{
					Id:               1,
					Name:             "Bitcoin",
					Tag:              "BTC",
					Algorithm:        "SHA-256",
					BlockTime:        "572.0",
					BlockReward:      6.34999333303702,
					BlockReward24:    6.35231784288925,
					BlockReward3:     6.350248762897884,
					BlockReward7:     6.335656042182499,
					LastBlock:        708163,
					Difficulty:       21659344833264,
					Difficulty24:     21659344833264,
					Difficulty3:      21659344833264,
					Difficulty7:      20896714582830.86,
					Nethash:          162633177817579450000,
					ExchangeRate:     61883.95,
					ExchangeRate24:   62466.1015202232,
					ExchangeRate3:    62350.343844704,
					ExchangeRate7:    61772.93474775763,
					ExchangeRateVol:  43852.82868,
					ExchangeRateCurr: "BTC",
					MarketCap:        "$1,167,338,979,536",
					PoolFee:          "0.000000",
					EstimatedRewards: "0.000413",
					BtcRevenue:       "0.00041284",
					Revenue:          "$25.55",
					Cost:             "$6.72",
					Profit:           "$18.83",
					Status:           "Active",
					Lagging:          false,
					Testing:          false,
					Listed:           true,
					Timestamp:        1636023038,
				},
			},
		},
		{
			name: "500",
			setup: func() {
				gock.New(URL).
					Get("coins/123.json").
					Reply(http.StatusInternalServerError)
			},
			input: input{
				id: 123,
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "403",
			setup: func() {
				gock.New(URL).
					Get("coins/123.json").
					Reply(http.StatusForbidden).
					JSON(map[string]interface{}{
						"errors": []string{"Could not find active coin with id 123"},
					})
			},
			input: input{
				id: 123,
			},
			output: output{
				withErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			actualRes, actualErr := API.GetCoinById(tt.input.id, "")
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}

func TestGetCoins(t *testing.T) {
	Cleanup(t)

	type output struct {
		res     Coins
		withErr bool
	}

	type test struct {
		name   string
		setup  func()
		output output
	}

	tests := []test{
		{
			name: "200",
			setup: func() {
				gock.New(URL).
					Get("coins.json").
					Reply(http.StatusOK).
					File(path.Join("json", "coins.json"))
			},
			output: output{
				res: Coins{
					Coins: CoinsMap{
						"Ethereum": Coin{
							Id:                 151,
							Tag:                "ETH",
							Algorithm:          "Ethash",
							BlockTime:          "13.5815",
							BlockReward:        2.24481498837968,
							BlockReward24:      2.22371827913691,
							LastBlock:          13550567,
							Difficulty:         10258652886885388,
							Difficulty24:       10417894079951800,
							Nethash:            755340197097919,
							ExchangeRate:       0.073304,
							ExchangeRate24:     0.0731101432584269,
							ExchangeRateVol:    5806.47714241,
							ExchangeRateCurr:   "BTC",
							MarketCap:          "$533,901,324,234.07",
							EstimatedRewards:   "0.0017",
							EstimatedRewards24: "0.00166",
							BtcRevenue:         "0.00012473",
							BtcRevenue24:       "0.00012167",
							Profitability:      100,
							Profitability24:    100,
							Lagging:            false,
							Timestamp:          1636033543,
						},
					},
				},
			},
		},
		{
			name: "500",
			setup: func() {
				gock.New(URL).
					Get("coins.json").
					Reply(http.StatusInternalServerError)
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "403",
			setup: func() {
				gock.New(URL).
					Get("coins.json").
					Reply(http.StatusForbidden)
			},
			output: output{
				withErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			actualRes, actualErr := API.GetCoins()
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}

func TestGetCalculators(t *testing.T) {
	Cleanup(t)
	type output struct {
		res     Calculators
		withErr bool
	}

	type test struct {
		name   string
		setup  func()
		output output
	}

	tests := []test{
		{
			name: "200",
			setup: func() {
				gock.New(URL).
					Get("calculators.json").
					Reply(http.StatusOK).
					File(path.Join("json", "calculators.json"))
			},
			output: output{
				res: Calculators{
					Coins: CalculatorsMap{
						"0xBitcoin": Calculator{
							Id:        315,
							Tag:       "0xBTC",
							Algorithm: "SHA3Solidity",
							Lagging:   true,
							Listed:    false,
							Status:    "Active",
							Testing:   false,
						},
						"365Coin": Calculator{
							Id:        74,
							Tag:       "365",
							Algorithm: "Keccak",
							Lagging:   true,
							Listed:    false,
							Status:    "No available stats",
							Testing:   false,
						},
					},
				},
			},
		},
		{
			name: "500",
			setup: func() {
				gock.New(URL).
					Get("calculators.json").
					Reply(http.StatusInternalServerError)
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "403",
			setup: func() {
				gock.New(URL).
					Get("calculators.json").
					Reply(http.StatusForbidden)
			},
			output: output{
				withErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			actualRes, actualErr := API.GetCalculators()
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}

func TestGet(t *testing.T) {
	Cleanup(t)
	type output struct {
		res     []byte
		withErr bool
	}
	type input struct {
		path string
	}
	type test struct {
		name   string
		setup  func()
		input  input
		output output
	}

	tests := []test{
		{
			name: "200",
			setup: func() {
				gock.New(URL).
					Get("coins/1.json").
					Reply(http.StatusOK).
					BodyString("Done")
			},
			input: input{
				path: "coins/1.json",
			},
			output: output{
				res: []byte("Done"),
			},
		},
		{
			name: "500",
			setup: func() {
				gock.New(URL).
					Get("calculators.json").
					Reply(http.StatusInternalServerError)
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "403",
			setup: func() {
				gock.New(URL).
					Get("calculators.json").
					Reply(http.StatusForbidden)
			},
			output: output{
				withErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			actualRes, actualErr := API.Get(tt.input.path)
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}
