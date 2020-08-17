package price

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type USDPrice struct {
	USD float64 `json:"USD"`
}

type Price struct {
	EthPrice USDPrice `json:"ETH"`
	TrxPrice USDPrice `json:"TRX"`
}

const (
	access_price_api string = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=ETH,TRX&tsyms=USD"
)

func GetPrice() (*Price, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(access_price_api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	p := &Price{}

	err = json.Unmarshal(body, p)
	if err != nil {
		return nil, err
	}

	return p, nil

}
