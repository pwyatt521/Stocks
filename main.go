package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pwyatt521/stocks/yahoo"
)

func main() {
	// Ask for a stock ticker
	fmt.Print("Enter a stock ticker: ")
	var symbol string
	_, err := fmt.Scanln(&symbol)
	if err != nil {
		log.Fatalf("Failed to read stock ticker: %v", err)
	}

	// Ask for start and end dates (optional)
	fmt.Println("Enter start date (YYYY-MM-DD), or leave empty to skip:")
	startDateInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	startDateInput = trimNewline(startDateInput)

	fmt.Println("Enter end date (YYYY-MM-DD), or leave empty to skip:")
	endDateInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	endDateInput = trimNewline(endDateInput)

	var startDate, endDate time.Time
	var useHistoricalData bool

	// Parse start and end dates if provided
	if startDateInput != "" && endDateInput != "" {
		startDate, err = time.Parse("2006-01-02", startDateInput)
		if err != nil {
			log.Fatalf("Failed to parse start date: %v", err)
		}

		endDate, err = time.Parse("2006-01-02", endDateInput)
		if err != nil {
			log.Fatalf("Failed to parse end date: %v", err)
		}

		useHistoricalData = true
	}

	// Retrieve stock data
	var stockData *yahoo.StockData
	var historicalPrices yahoo.YahooFinanceHistoricalResponse
	var apiErr error

	if useHistoricalData {
		historicalPrices, apiErr = yahoo.GetStockHistoricalPrices(symbol, startDate, endDate)
	} else {
		stockData, apiErr = yahoo.GetStockData(symbol)
	}

	if apiErr != nil {
		log.Fatalf("Failed to retrieve stock data: %v", apiErr)
	}

	// Print the results
	fmt.Println("Stock Data:")
	if useHistoricalData {
		historicalPrices.Print()
	} else {
		fmt.Printf("Symbol: %s\n", stockData.Symbol)
		fmt.Printf("Current Price: %.2f\n", stockData.CurrentPrice)
	}
}

func trimNewline(s string) string {
	return s[:len(s)-1]
}
