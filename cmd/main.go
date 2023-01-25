package main

import (
	"context"
	"flag"
	"fmt"
	cli "github.com/vijaynallagatla/cli-stocks/yapi"
	"net/http"
)

func main() {
	cfg := cli.NewConfiguration()
	client := cli.NewAPIClient(cfg)

	// Parse symbols through cli args -s for symbol and use .NS or . for changing the stock exchange
	symbol := flag.String("s", "AAPL", "Stock Symbol")
	flag.Parse()

	// Ignoring err since there could be marshalling issue with some response object
	res, httpStatus, _ := client.QuoteApi.GetQuote(context.Background()).Symbols(*symbol).Execute()

	if httpStatus != nil {
		if httpStatus.StatusCode != http.StatusOK {
			panic(httpStatus)
		}
	}

	quote, ok := res.GetQuoteResponseOk()
	if !ok {
		panic("error getting response")
	}

	for _, stock := range quote.GetResult() {
		fmt.Printf(" Stock Name: %v \n Current Value: %v\n", stock.GetShortName(), stock.GetRegularMarketPrice())
	}
}
