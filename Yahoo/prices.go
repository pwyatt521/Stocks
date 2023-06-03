package yahoo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)


type YahooFinanceHistoricalResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency                 string `json:"currency"`
				Symbol                   string `json:"symbol"`
				ExchangeName             string `json:"exchangeName"`
				InstrumentType           string `json:"instrumentType"`
				FirstTradeDate           int64  `json:"firstTradeDate"`
				RegularMarketTime        int64  `json:"regularMarketTime"`
				GMTOffset                int    `json:"gmtoffset"`
				Timezone                 string `json:"timezone"`
				ExchangeTimezoneName     string `json:"exchangeTimezoneName"`
				RegularMarketPrice       float64 `json:"regularMarketPrice"`
				ChartPreviousClose       float64 `json:"chartPreviousClose"`
				PriceHint                int     `json:"priceHint"`
				DataGranularity          string  `json:"dataGranularity"`
				Range                    string  `json:"range"`
				ValidRanges              []string `json:"validRanges"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int64  `json:"start"`
						End       int64  `json:"end"`
						GMTOffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int64  `json:"start"`
						End       int64  `json:"end"`
						GMTOffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int64  `json:"start"`
						End       int64  `json:"end"`
						GMTOffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
			} `json:"meta"`
			Timestamp []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []float64 `json:"close"`
					Low  []float64 `json:"low"`
					Open []float64 `json:"open"`
					High []float64 `json:"high"`
					Volume []float64 `json:"volume"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

func (r *YahooFinanceHistoricalResponse) Print() {
	fmt.Println("Currency:", r.Chart.Result[0].Meta.Currency)
	fmt.Println("Symbol:", r.Chart.Result[0].Meta.Symbol)
	fmt.Println("Exchange Name:", r.Chart.Result[0].Meta.ExchangeName)
	fmt.Println("Instrument Type:", r.Chart.Result[0].Meta.InstrumentType)
	fmt.Println("First Trade Date:", FormatInt64ToDate(r.Chart.Result[0].Meta.FirstTradeDate))
	fmt.Println("Regular Market Time:", FormatInt64ToTime(r.Chart.Result[0].Meta.RegularMarketTime))
	fmt.Println("GMT Offset:", r.Chart.Result[0].Meta.GMTOffset)
	fmt.Println("Timezone:", r.Chart.Result[0].Meta.Timezone)
	fmt.Println("Exchange Timezone Name:", r.Chart.Result[0].Meta.ExchangeTimezoneName)
	fmt.Println("Regular Market Price:", r.Chart.Result[0].Meta.RegularMarketPrice)
	fmt.Println("Chart Previous Close:", r.Chart.Result[0].Meta.ChartPreviousClose)
	fmt.Println("Price Hint:", r.Chart.Result[0].Meta.PriceHint)
	fmt.Println("Data Granularity:", r.Chart.Result[0].Meta.DataGranularity)
	fmt.Println("Range:", r.Chart.Result[0].Meta.Range)
	fmt.Println("Valid Ranges:", r.Chart.Result[0].Meta.ValidRanges)
	fmt.Print("Timestamp: ")
	for _, tStamp := range r.Chart.Result[0].Timestamp {
		fmt.Print(FormatInt64ToTime(tStamp)+ ", ")
	}
	fmt.Println("")

	for i := 0; i < len(r.Chart.Result[0].Indicators.Quote); i++ {
		fmt.Println("Low:", r.Chart.Result[0].Indicators.Quote[i].Low)
		fmt.Println("Open:", r.Chart.Result[0].Indicators.Quote[i].Open)
		fmt.Println("High:", r.Chart.Result[0].Indicators.Quote[i].High)
		fmt.Println("Close:", r.Chart.Result[0].Indicators.Quote[i].Close)
		fmt.Println("Volume:", r.Chart.Result[0].Indicators.Quote[i].Volume)
	}
}

func GetStockHistoricalPrices(symbol string, startDate, endDate time.Time)(YahooFinanceHistoricalResponse, error) {
	// Convert start and end dates to Unix timestamps
	startUnix := startDate.Unix()
	endUnix := endDate.Unix()

	// Make API call to retrieve historical stock prices
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/chart/%s?period1=%d&period2=%d&interval=1d", symbol, startUnix, endUnix)
	response, err := http.Get(url)
	if err != nil {
		return YahooFinanceHistoricalResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return YahooFinanceHistoricalResponse{}, err

	}

	// Parse the API response
	var yahooResponse YahooFinanceHistoricalResponse
	err = json.Unmarshal(body,&yahooResponse)
	if err != nil {
		return YahooFinanceHistoricalResponse{}, err
	}

	return yahooResponse, nil
}
