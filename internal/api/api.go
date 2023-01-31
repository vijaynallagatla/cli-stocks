package api

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	cli "github.com/vijaynallagatla/cli-stocks/yapi"
	"net/http"
	"os"
)

func GetQuoteForSymbol(client *cli.APIClient, symbol *string) error {
	quote := quoteForSymbol(client, symbol)
	columns := []table.Column{
		{Title: "Stock Name", Width: 30},
		{Title: "Stock Value", Width: 20},
		{Title: "Previous Close", Width: 20},
		{Title: "Day Range", Width: 20},
	}

	rows := []table.Row{
		{quote.name, quote.currentValue, fmt.Sprintf("%v", quote.previousDayClose), quote.regularMarketDayRange},
	}

	// Send the UI for rendering
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)),
		table.WithStyles(getTableStyles()),
	)

	p := tea.NewProgram(initialModel(quote, t))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	return nil
}

func quoteForSymbol(client *cli.APIClient, symbol *string) stock {
	// Ignoring err since there could be marshalling issue with some response object
	res, httpStatus, _ := client.QuoteApi.GetQuote(context.Background()).Symbols(*symbol).Execute()
	if httpStatus != nil {
		if httpStatus.StatusCode != http.StatusOK {
			panic(httpStatus)
		}
	}

	quote, ok := res.GetQuoteResponseOk()
	if !ok {
		panic("error getting response from yahoo finance")
	}

	if len(quote.GetResult()) > 0 {
		return stock{
			name:                  quote.GetResult()[0].GetShortName(),
			currentValue:          fmt.Sprintf("%f", quote.GetResult()[0].GetRegularMarketPrice()),
			regularMarketDayRange: quote.GetResult()[0].GetRegularMarketDayRange(),
			previousDayClose:      quote.GetResult()[0].GetRegularMarketPreviousClose(),
		}
	}

	return stock{
		name: "Not Found",
	}
}
