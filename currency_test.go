package currency_converter

import (
	gock "gopkg.in/h2non/gock.v1"

	"net/http"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nexeranet/currency_converter/pkg/coingecko"
	"github.com/nexeranet/currency_converter/pkg/whattomine"
)

var API = NewConverter()
var COINGECKOAPI = "https://api.coingecko.com/api/v3"
var WHATTOMINEAPI = "https://whattomine.com"
var setUpIsDone = false

func RunSetup() {
	if !setUpIsDone {
		API.Setup()
		setUpIsDone = true
	}
}

func Cleanup(t *testing.T) {
	t.Cleanup(func() {
		gock.EnableNetworking()
		gock.OffAll()
	})
	gock.DisableNetworking()

	RunSetup()
}

func TestGetNetInfo(t *testing.T) {
	Cleanup(t)

	type input struct {
		tag string
	}

	type output struct {
		res     whattomine.Coin
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
				gock.New(WHATTOMINEAPI).
					Get("coins/1.json").
					Reply(http.StatusOK).
					File(path.Join("json", "netinfo.json"))
			},
			input: input{
				tag: "BTC",
			},
			output: output{
				res: whattomine.Coin{
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
				gock.New(WHATTOMINEAPI).
					Get("coins/123.json").
					Reply(http.StatusInternalServerError)
			},
			input: input{
				tag: "BTC@#",
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "403",
			setup: func() {
				gock.New(WHATTOMINEAPI).
					Get("coins/123.json").
					Reply(http.StatusForbidden).
					JSON(map[string]interface{}{
						"errors": []string{"Could not find active coin with id 123"},
					})
			},
			input: input{
				tag: "BTC@#",
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
			actualRes, actualErr := API.GetNetInfo(tt.input.tag)
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}

func TestGetPrice(t *testing.T) {
	Cleanup(t)

	type input struct {
		currency  string
		convertTo string
	}

	type output struct {
		res     float32
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
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusOK).
					File(path.Join("json", "simple_price.json"))
			},
			input: input{
				currency:  "BTC",
				convertTo: "USD",
			},
			output: output{
				res: 61510,
			},
		},
		{
			name: "500",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusInternalServerError)
			},
			input: input{},
			output: output{
				withErr: true,
			},
		},
		{
			name: "400",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusBadRequest)
			},
			input: input{},
			output: output{
				withErr: true,
			},
		},
		{
			name: "Invalid currency tag",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusBadRequest)
			},
			input: input{
				currency:  "BTCDFDD",
				convertTo: "ETH",
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "Invalid convertTo currency tag",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusBadRequest)
			},
			input: input{
				currency:  "BTC",
				convertTo: "ETH!@#@",
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
			actualRes, actualErr := API.GetPrice(tt.input.currency, tt.input.convertTo)
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}

func TestGetPricesInGroups(t *testing.T) {
	Cleanup(t)

	type input struct {
		currency  []string
		convertTo []string
	}

	type output struct {
		res     coingecko.PricesMap
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
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusOK).
					File(path.Join("json", "simple_price_in_group.json"))
			},
			input: input{
				currency:  []string{"BTC", "ETH"},
				convertTo: []string{"USD"},
			},
			output: output{
				res: coingecko.PricesMap{
					"BTC": {"USD": 5005.73},
					"ETH": {"USD": 163.58},
				},
			},
		},
		{
			name: "500",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusInternalServerError)
			},
			input: input{},
			output: output{
				withErr: true,
			},
		},
		{
			name: "400",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusBadRequest)
			},
			input: input{},
			output: output{
				withErr: true,
			},
		},
		{
			name: "Invalid currency tag",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusBadRequest)
			},
			input: input{
				currency:  []string{"BTCDFDD"},
				convertTo: []string{"ETH"},
			},
			output: output{
				withErr: true,
			},
		},
		{
			name: "Invalid convertTo currency tag",
			setup: func() {
				gock.New(COINGECKOAPI).
					Get("/simple/price").
					Reply(http.StatusBadRequest)
			},
			input: input{
				currency:  []string{"BTC"},
				convertTo: []string{"ETH!@#@"},
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
			actualRes, actualErr := API.GetPricesInGroups(tt.input.currency, tt.input.convertTo)
			if (actualErr != nil) != tt.output.withErr {
				t.Fatalf("expected error %t, actual %s", tt.output.withErr, actualErr)
			}
			if !cmp.Equal(tt.output.res, actualRes) {
				t.Fatalf("expected output do not match\n%s", cmp.Diff(tt.output.res, actualRes))
			}
		})
	}
}
