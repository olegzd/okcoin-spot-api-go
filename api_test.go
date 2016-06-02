package okcoin

import (
	"os"
	"testing"
)

func TestSpotPrice(t *testing.T) {

	// Test that getSpotPrice correctly calls
	var sp = SpotPrice{}

	getSpotPrice("ltc_usd", &sp)

	t.Logf("Ticker date: %s\n", sp.Date)
	t.Logf("Ticker high: %s\n", sp.Ticker.High)
	t.Logf("Ticker low: %s\n", sp.Ticker.Low)
	t.Logf("Ticker volume: %s\n", sp.Ticker.Vol)
}

func TestAccountInfo(t *testing.T) {
	var apiKey = os.Getenv("OKCOIN_API_KEY")
	var apiSecretKey = os.Getenv("OKCOIN_SECRET_KEY")

	var account = Account{apiKey, apiSecretKey}
	// var keys = map[string]string{"api_key": account.APIKey}

	// Get account info
	info := Info{}
	account.getAccountInfo(&info)

}
