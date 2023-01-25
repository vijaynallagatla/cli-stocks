package main

import (
	"context"
	cli "github.com/vijaynallagatla/cli-stocks/client"
	"net/http"
)

func main() {
	cfg := cli.NewConfiguration()
	client := cli.NewAPIClient(cfg)

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
