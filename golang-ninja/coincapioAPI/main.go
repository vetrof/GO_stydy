package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const sampleUrl = "http://api.coincap.io/v2/assets"

type assetData struct {
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

type assetsResponse struct {
	Data      []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type LoggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (d assetData) Info() string {
	return fmt.Sprintf("[ID]: %s | [Rank]: %s | [Name]: %s", d.ID, d.Rank, d.Name)
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, ">>> LOGGER: [%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("REDIRECT", req.Response.Status)
			return nil
		},
		Transport: &LoggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(sampleUrl)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf(">>>>>>>>>>>>>>> STATUS %d <<<<<<<<<<<<<<<<", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf(string(body))

	var r assetsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range r.Data {
		fmt.Println(asset.Info())
	}
}
