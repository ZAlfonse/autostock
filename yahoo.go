package autostock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//https://query2.finance.yahoo.com/v10/finance/quoteSummary/AMD?lang=en-US&region=US&modules=upgradeDowngradeHistory%2CrecommendationTrend
const (
	YahooScheme      = "https"
	YahooDomain1     = "query1.finance.yahoo.com"
	YahooDomain2     = "query2.finance.yahoo.com"
	YahooAPIVersion  = "v10"
	YahooAPIRoot     = "finance/quoteSummary"
	YahooAlwaysArgs  = "lang=en-US&region=US"
	YahooReccModules = "modules=upgradeDowngradeHistory%2CrecommendationTrend"
)

//Trend ...
type Trend struct {
	Period     string `json:"period"`
	StrongBuy  int    `json:"strongBuy"`
	Buy        int    `json:"buy"`
	Hold       int    `json:"hold"`
	Sell       int    `json:"sell"`
	StrongSell int    `json:"strongSell"`
}

//History ...
type History struct {
	EpochGradeDate int    `json:"epochGradeDate"`
	Firm           string `json:"firm"`
	ToGrade        string `json:"toGrade"`
	FromGrade      string `json:"fromGrade"`
	Action         string `json:"action"`
}

//UpgradeDowngradeHistory ...
type UpgradeDowngradeHistory struct {
	History []History `json:"history"`
	maxAge  int
}

//RecommendationTrend ...
type RecommendationTrend struct {
	Trend  []Trend `json:"trend"`
	maxAge int
}

//RecommendationResults ...
type RecommendationResults struct {
	RecommendationTrend     RecommendationTrend     `json:"recommendationTrend"`
	UpgradeDowngradeHistory UpgradeDowngradeHistory `json:"upgradeDowngradeHistory"`
}

//QuoteSummary ...
type QuoteSummary struct {
	Result []RecommendationResults `json:"result"`
	Error  string                  `json:"error"`
}

//RecommendationResponse ...
type RecommendationResponse struct {
	QuoteSummary QuoteSummary `json:"quoteSummary"`
}

//YahooAnalystInfo ...
func YahooAnalystInfo(ticker string) []History {
	u := new(url.URL)
	u.Scheme = YahooScheme
	u.Host = YahooDomain1
	u.Path = fmt.Sprintf("/%v/%v/%v", YahooAPIVersion, YahooAPIRoot, ticker)
	u.RawQuery = fmt.Sprintf("%v&%v", YahooAlwaysArgs, YahooReccModules)

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(u)
		fmt.Println(resp)
		fmt.Println(err)
		panic("Problem requesting Yahoo API")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body from request")
	}

	byt := []byte(body)

	dat := RecommendationResponse{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		fmt.Print(body)
		panic(err)
	}

	return dat.QuoteSummary.Result[0].UpgradeDowngradeHistory.History
}
