package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	cli "github.com/vijaynallagatla/cli-stocks/yapi"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	stock    stock
	table    table.Model
	keymap   keymap
	help     help.Model
	quitting bool
}

type stock struct {
	name         string
	currentValue string
}
type keymap struct {
	stop key.Binding
	quit key.Binding
}

func initialModel(stock stock, t table.Model) model {
	return model{
		stock:    stock,
		table:    t,
		keymap:   keymap{},
		help:     help.Model{},
		quitting: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q", "enter", "stop":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.stop,
	})
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\nPress 'q' to Quit!\n"
}

func getTableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return s
}

func getQuoteForSymbol(client *cli.APIClient, symbol *string) stock {
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

	if len(quote.GetResult()) > 0 {
		return stock{
			name:         quote.GetResult()[0].GetShortName(),
			currentValue: fmt.Sprintf("%f", quote.GetResult()[0].GetRegularMarketPrice()),
		}
	}

	return stock{
		name: "Not Found",
	}
}

func main() {
	cfg := cli.NewConfiguration()
	client := cli.NewAPIClient(cfg)

	// Parse symbols through cli args -s for symbol and use .NS or . for changing the stock exchange
	symbol := flag.String("s", "AAPL", "Stock Symbol")
	flag.Parse()

	quote := getQuoteForSymbol(client, symbol)

	columns := []table.Column{
		{Title: "Stock Name", Width: 30},
		{Title: "Stock Value", Width: 20},
	}

	rows := []table.Row{
		{quote.name, quote.currentValue},
	}

	// Send the UI for rendering
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
		table.WithStyles(getTableStyles()),
	)

	p := tea.NewProgram(initialModel(quote, t))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
