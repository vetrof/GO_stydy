package coincap

import "fmt"

type AssetsResponse struct {
	Data      []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type AssetResponse struct {
	Data      Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID]: %s | [Rank]: %s | [Name]: %s", d.ID, d.Rank, d.Name)
}
