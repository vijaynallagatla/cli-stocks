package api

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
)

type model struct {
	stock    stock
	table    table.Model
	keymap   keymap
	help     help.Model
	quitting bool
}

type stock struct {
	name                  string
	currentValue          string
	currency              string
	previousDayClose      float32
	regularMarketDayRange string
}
type keymap struct {
	stop key.Binding
	quit key.Binding
}

type search struct {
	// Exact Symbol match search for quotation
	symbol *string
	name   *string
}

type Track struct {
	command   string
	Flag      string
	FlagValue string
}

var tracks = []Track{
	{"quote", "--s", "AAPL"},
	{"quote", "--s", "GOOG"},
}
