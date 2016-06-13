/* Package spotapi provides OKCoin's Spot trading/price API.
 * More of the documentation can be found here:
 * https://www.okcoin.com/about/rest_api.do
 */

package spotapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	api = "https://www.okcoin.com/api/v1/"
)

// AccountInfo is the main struct containing response info from API
type AccountInfo struct {
	Info   Info `json:"info"`
	Result bool `json:"result"`
}

// Account contains the keys necessary to access the OKCoin API
type Account struct {
	APIKey    string
	SecretKey string
}

// Info is the account info main struct
type Info struct {
	Fund Funds `json:"funds"`
}

// Funds encompasses assets and amounts of the account
type Funds struct {
	Asset     Asset  `json:"asset"`
	Borrow    Amount `json:"borrow"`
	Free      Amount `json:"free"`
	Freezed   Amount `json:"freezed"`
	UnionFund Amount `json:"union_fund"`
}

// Amount struct contains how much btc, ltc, and USD
// in your account for a specific fund
type Amount struct {
	BTC string `json:"btc"`
	USD string `json:"usd"`
	LTC string `json:"ltc"`
}

// Asset describes account asset total worth
type Asset struct {
	Net   string `json:"net"`
	Total string `json:"total"`
}

// SpotPrice is used to receive the latest OKCoin spot market data
type SpotPrice struct {
	// Date is server time for returned data
	Date string `json:"date"`

	// Ticker contains spot price data
	Ticker `json:"ticker"`
}

// Ticker is a struct used to get spot price
type Ticker struct {

	// Buy is the best bid
	Buy string `json:"buy"`

	// High is the highest price
	High string `json:"high"`

	// Last is latest price
	Last string `json:"last"`

	// Low is lowest price
	Low string `json:"low"`

	// Sell is the best ask
	Sell string `json:"sell"`

	// Vol is volume in the last rolling 24 hours
	Vol string `json:"vol"`
}

// getSpotPrice gets the current spot price from API
// depending on the symbol. symbol can either be 'ltc_usd' or
// 'btc_usd' ; anything else results in btc_usd
func getSpotPrice(symbol string, target *SpotPrice) error {
	url := api + "ticker.do?symbol=" + symbol
	resp, err := http.Get(url)
	defer resp.Body.Close()

	// check for errors, and unmarshal the response to Ticker struct
	if err != nil {
		return err
	}

	json.NewDecoder(resp.Body).Decode(target)
	return nil
}

// getAccountInfo fetches account info and populates target with result
func (a *Account) getAccountInfo(target *AccountInfo) {

	// Generate sign and prepare POST params
	signature := generateSign(map[string]string{"api_key": a.APIKey}, a.SecretKey)
	values := url.Values{}

	values.Add("api_key", a.APIKey)
	values.Add("sign", signature)

	response, error := http.PostForm("https://www.okcoin.com/api/v1/userinfo.do", values)

	// Check for errors and panic if one arose, otherwise populate target
	if error != nil {
		panic(error)
	}
	json.NewDecoder(response.Body).Decode(target)
}

// generateSign takes in a set of params (map of string:string) and
// returns an MD5 hash signature
func generateSign(params map[string]string, secretKey string) (signature string) {
	// Get the keys into an array, then sort it
	var keys = make([]string, len(params))

	i := 0
	for k := range params {
		keys[i] = k
		i++
	}

	// Sort keys and append each param to the signature, seperated by '&'
	sort.Strings(keys)
	for _, key := range keys {
		signature = signature + key + "=" + params[key] + "&"
	}

	// Add the private key to the signature
	signature += "secret_key=" + secretKey

	// MD5 Hash the signature
	hasher := md5.New()
	hasher.Write([]byte(signature))
	signature = strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))

	return signature
}
