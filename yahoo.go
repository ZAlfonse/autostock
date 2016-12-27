package autostock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//https://query2.finance.yahoo.com/v10/finance/quoteSummary/AMD?lang=en-US&region=US&modules=upgradeDowngradeHistory%2CrecommendationTrend
const (
	YahooScheme     = "https"
	YahooDomain1    = "query1.finance.yahoo.com"
	YahooDomain2    = "query2.finance.yahoo.com"
	YahooAPIVersion = "v10"
	YahooAPIRoot    = "finance/quoteSummary"
	YahooAlwaysArgs = "lang=en-US&region=US"

	YahooRatingHistoryModule = "upgradeDowngradeHistory"
	YahooRatingTrendModule   = "recommendationTrend"
	YahooFinancialDataModule = "financialData"

	YahooMaxRecommendationRating = 5
)

//YahooMakeRequest ...
func YahooMakeRequest(ticker string, modules []string) YahooResult {
	u := new(url.URL)
	u.Scheme = YahooScheme
	u.Host = YahooDomain1
	u.Path = fmt.Sprintf("/%v/%v/%v", YahooAPIVersion, YahooAPIRoot, ticker)
	u.RawQuery = fmt.Sprintf("%v&modules=%v", YahooAlwaysArgs, strings.Join(modules, "%2C"))

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(u)
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(u)
		panic(err)
	}

	byt := []byte(body)

	yResp := YahooResponse{}

	if err := json.Unmarshal(byt, &yResp); err != nil {
		fmt.Println(u)
		fmt.Println(err)
	}

	return yResp.QuoteSummary.Result[0]
}
