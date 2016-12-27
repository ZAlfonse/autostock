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
func BulkRecommendSummary(tickers []string, onlytrue bool) {
	fmt.Println("Ticker  CurrentPrice  MeanTargetPrice  LowTargetPrice  HighTargetPrice  AnalystConsensus  NumAnalysts  Recommend?")
	for _, ticker := range tickers {
		result := YahooMakeRequest(ticker, []string{
			YahooFinancialDataModule,
		})

		currentPrice := result.FinancialData.CurrentPrice.Raw
		lowTargetPrice := result.FinancialData.TargetLowPrice.Raw
		meanTargetPrice := result.FinancialData.TargetMeanPrice.Raw
		highTargetPrice := result.FinancialData.TargetHighPrice.Raw
		currentAnalystConsensus := result.FinancialData.RecommendationMean.Raw
		totalAnalysts := result.FinancialData.NumberOfAnalystOpinions.Raw

		recommend := currentPrice < meanTargetPrice && currentAnalystConsensus > 3 && totalAnalysts > 2 && highTargetPrice-lowTargetPrice < meanTargetPrice/1.5

		if onlytrue {
			if recommend {
				fmt.Printf("%-8v%-14v%-17v%-16v%-17v%-18v%-15v%v\n", ticker, currentPrice, meanTargetPrice, lowTargetPrice, highTargetPrice, currentAnalystConsensus, totalAnalysts, recommend)
			}
		} else {
			fmt.Printf("%-8v%-14v%-17v%-16v%-17v%-18v%-15v%v\n", ticker, currentPrice, meanTargetPrice, lowTargetPrice, highTargetPrice, currentAnalystConsensus, totalAnalysts, recommend)
		}

	}
}
