package yahoo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StockData struct {
	Symbol           string    `json:"symbol"`
	CurrentPrice     float64   `json:"current_price"`
	HistoricalPrices []float64 `json:"historical_prices"`
}

type yahooFinanceResponse struct {
	Price struct {
		RegularMarketPrice float64 `json:"regularMarketPrice"`
	} `json:"price"`
	Chart struct {
		Result []struct {
			Timestamp  []int64   `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

func GetStockData(symbol string) (*StockData, error) {
	// Make API call to retrieve stock data
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the API response
	var yahooResponse yahooFinanceResponse
	err = json.NewDecoder(response.Body).Decode(&yahooResponse)
	if err != nil {
		return nil, err
	}

	// Extract relevant data from the response
	var historicalPrices []float64
	for _, closePrice := range yahooResponse.Chart.Result[0].Indicators.Quote[0].Close {
		historicalPrices = append(historicalPrices, closePrice)
	}

	stockData := &StockData{
		Symbol:           symbol,
		CurrentPrice:     yahooResponse.Price.RegularMarketPrice,
		HistoricalPrices: historicalPrices,
	}

	return stockData, nil
}

