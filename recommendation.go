package autostock

import "fmt"

//Recommend ...
func Recommend(ticker string) bool {
	result := YahooMakeRequest(ticker, []string{
		YahooFinancialDataModule,
	})

	currentPrice := result.FinancialData.CurrentPrice.Raw
	meanTargetPrice := result.FinancialData.TargetMeanPrice.Raw
	currentAnalystConsensus := result.FinancialData.RecommendationMean.Raw

	recommend := currentPrice < meanTargetPrice && currentAnalystConsensus > 3

	return recommend
}

//BulkRecommendSummary ...
func BulkRecommendSummary(tickers []string) {
	fmt.Println("Ticker  CurrentPrice  MeanTargetPrice  AnalystConsensus  Recommend?")
	for _, ticker := range tickers {
		result := YahooMakeRequest(ticker, []string{
			YahooFinancialDataModule,
		})

		currentPrice := result.FinancialData.CurrentPrice.Raw
		meanTargetPrice := result.FinancialData.TargetMeanPrice.Raw
		currentAnalystConsensus := result.FinancialData.RecommendationMean.Raw

		recommend := currentPrice < meanTargetPrice && currentAnalystConsensus > 3

		fmt.Printf("%-8v%-14v%-17v%-18v%v\n", ticker, currentPrice, meanTargetPrice, currentAnalystConsensus, recommend)
	}
}
