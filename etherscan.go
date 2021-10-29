package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ncruces/zenity"
)

var (
	apiKey string
	apiUrl = getAPIURL()
)

type Result struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		// LastBlock       string `json:"LastBlock"`
		SafeGasPrice string `json:"SafeGasPrice"`
		// ProposeGasPrice string `json:"ProposeGasPrice"`
		// FastGasPrice    string `json:"FastGasPrice"`
		// SuggestBaseFee  string `json:"suggestBaseFee"`
		// GasUsedRatio    string `json:"gasUsedRatio"`
	} `json:"result"`
}

func getAPIURL() string {
	if apiKey == "" {
		apiKey, _ = zenity.Entry("Enter your Etherscan Api Key\nSee: https://docs.etherscan.io/getting-started/viewing-api-usage-statistics",
			zenity.Title("Enter api key"))
	}
	return fmt.Sprintf("https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=%s", apiKey)
}

func getGwei() string {
	r := &Result{}
	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	return r.Result.SafeGasPrice
}
