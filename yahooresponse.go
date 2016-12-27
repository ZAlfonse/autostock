package autostock

//YahooResult ...
type YahooResult struct {
	EarningsHistory struct {
		History []struct {
			EpsActual struct {
				Fmt string  `json:"fmt"`
				Raw float64 `json:"raw"`
			} `json:"epsActual"`
			EpsDifference struct {
				Fmt string `json:"fmt"`
				Raw int64  `json:"raw"`
			} `json:"epsDifference"`
			EpsEstimate struct {
				Fmt string  `json:"fmt"`
				Raw float64 `json:"raw"`
			} `json:"epsEstimate"`
			MaxAge  int64  `json:"maxAge"`
			Period  string `json:"period"`
			Quarter struct {
				Fmt string `json:"fmt"`
				Raw int64  `json:"raw"`
			} `json:"quarter"`
			SurprisePercent struct {
				Fmt string `json:"fmt"`
				Raw int64  `json:"raw"`
			} `json:"surprisePercent"`
		} `json:"history"`
		MaxAge int64 `json:"maxAge"`
	} `json:"earningsHistory"`
	EarningsTrend struct {
		MaxAge int64 `json:"maxAge"`
		Trend  []struct {
			EarningsEstimate struct {
				Avg struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"avg"`
				Growth struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"growth"`
				High struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"high"`
				Low struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"low"`
				NumberOfAnalysts struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"numberOfAnalysts"`
				YearAgoEps struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"yearAgoEps"`
			} `json:"earningsEstimate"`
			EndDate      string `json:"endDate"`
			EpsRevisions struct {
				DownLast30days struct {
					Fmt     interface{} `json:"fmt"`
					LongFmt string      `json:"longFmt"`
					Raw     int64       `json:"raw"`
				} `json:"downLast30days"`
				DownLast90days struct{} `json:"downLast90days"`
				UpLast30days   struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"upLast30days"`
				UpLast7days struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"upLast7days"`
			} `json:"epsRevisions"`
			EpsTrend struct {
				Three0daysAgo struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"30daysAgo"`
				Six0daysAgo struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"60daysAgo"`
				SevenDaysAgo struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"7daysAgo"`
				Nine0daysAgo struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"90daysAgo"`
				Current struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"current"`
			} `json:"epsTrend"`
			Growth struct {
				Fmt string  `json:"fmt"`
				Raw float64 `json:"raw"`
			} `json:"growth"`
			MaxAge          int64  `json:"maxAge"`
			Period          string `json:"period"`
			RevenueEstimate struct {
				Avg struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"avg"`
				Growth struct {
					Fmt string  `json:"fmt"`
					Raw float64 `json:"raw"`
				} `json:"growth"`
				High struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"high"`
				Low struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"low"`
				NumberOfAnalysts struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"numberOfAnalysts"`
				YearAgoRevenue struct {
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
					Raw     int64  `json:"raw"`
				} `json:"yearAgoRevenue"`
			} `json:"revenueEstimate"`
		} `json:"trend"`
	} `json:"earningsTrend"`
	FinancialData struct {
		CurrentPrice struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"currentPrice"`
		CurrentRatio struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"currentRatio"`
		DebtToEquity struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"debtToEquity"`
		EarningsGrowth struct{} `json:"earningsGrowth"`
		Ebitda         struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"ebitda"`
		EbitdaMargins struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"ebitdaMargins"`
		FreeCashflow struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"freeCashflow"`
		GrossMargins struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"grossMargins"`
		GrossProfits struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"grossProfits"`
		MaxAge                  int64 `json:"maxAge"`
		NumberOfAnalystOpinions struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"numberOfAnalystOpinions"`
		OperatingCashflow struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"operatingCashflow"`
		OperatingMargins struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"operatingMargins"`
		ProfitMargins struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"profitMargins"`
		QuickRatio struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"quickRatio"`
		RecommendationKey  string `json:"recommendationKey"`
		RecommendationMean struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"recommendationMean"`
		ReturnOnAssets struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"returnOnAssets"`
		ReturnOnEquity struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"returnOnEquity"`
		RevenueGrowth struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"revenueGrowth"`
		RevenuePerShare struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"revenuePerShare"`
		TargetHighPrice struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"targetHighPrice"`
		TargetLowPrice struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"targetLowPrice"`
		TargetMeanPrice struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"targetMeanPrice"`
		TargetMedianPrice struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"targetMedianPrice"`
		TotalCash struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"totalCash"`
		TotalCashPerShare struct {
			Fmt string  `json:"fmt"`
			Raw float64 `json:"raw"`
		} `json:"totalCashPerShare"`
		TotalDebt struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"totalDebt"`
		TotalRevenue struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"totalRevenue"`
	} `json:"financialData"`
	IndustryTrend struct {
		Estimates []interface{} `json:"estimates"`
		MaxAge    int64         `json:"maxAge"`
		PeRatio   struct{}      `json:"peRatio"`
		PegRatio  struct{}      `json:"pegRatio"`
		Symbol    interface{}   `json:"symbol"`
	} `json:"industryTrend"`
	RecommendationTrend struct {
		MaxAge int64 `json:"maxAge"`
		Trend  []struct {
			Buy        int64  `json:"buy"`
			Hold       int64  `json:"hold"`
			Period     string `json:"period"`
			Sell       int64  `json:"sell"`
			StrongBuy  int64  `json:"strongBuy"`
			StrongSell int64  `json:"strongSell"`
		} `json:"trend"`
	} `json:"recommendationTrend"`
	UpgradeDowngradeHistory struct {
		History []struct {
			Action         string `json:"action"`
			EpochGradeDate int64  `json:"epochGradeDate"`
			Firm           string `json:"firm"`
			FromGrade      string `json:"fromGrade"`
			ToGrade        string `json:"toGrade"`
		} `json:"history"`
		MaxAge int64 `json:"maxAge"`
	} `json:"upgradeDowngradeHistory"`
}

//YahooResponse ...
type YahooResponse struct {
	QuoteSummary struct {
		Error  interface{}   `json:"error"`
		Result []YahooResult `json:"result"`
	} `json:"quoteSummary"`
}
