package spotapi

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

	// Get account info (populate via Info struct)
	info := AccountInfo{}
	account.getAccountInfo(&info)

	// Print out account info
	t.Logf("Result status: %t\n", info.Result)
	t.Logf("info.Info.Fund.Asset.Total: %s\n", info.Info.Fund.Asset.Total)
	t.Logf("info.Info.Fund.Borrow.USD: %s\n", info.Info.Fund.Borrow.USD)
	t.Logf("info.Info.Fund.Free.USD: %s\n", info.Info.Fund.Free.USD)
	//... and etc.
}
