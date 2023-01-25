package main

import (
	"context"
	"github.com/vijaynallagatla/cli_stocks/client"
	"net/http"
)

func main() {
	cfg := yapi.NewConfiguration()
	client := yapi.NewAPIClient(cfg)

	res, httpStatus, err := client.QuoteApi.GetQuote(context.Background()).Symbols("AAPL").Execute()
	if err != nil {
		panic(err)
	}

	if httpStatus != nil {
		if httpStatus.StatusCode != http.StatusOK {
			panic(httpStatus)
		}
	}

	r, ok := res.GetQuoteResponseOk()
	if !ok {
		panic(err)
	}

	if r != nil {
		panic(r)
	}
}
