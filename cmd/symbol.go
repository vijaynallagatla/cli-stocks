/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vijaynallagatla/cli-stocks/internal/api"
)

// symbolCmd represents the symbol command
var symbolCmd = &cobra.Command{
	Use:              "quote [no options!]",
	TraverseChildren: true,
	Short:            "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("s")
		if len(symbol) > 0 {
			api.GetQuoteForSymbol(apiClient, &symbol)
		} else {
			fmt.Println("Please provide the symbol")
		}
	},
}

func init() {
	rootCmd.AddCommand(symbolCmd)
	symbolCmd.PersistentFlags().String("s", "", "A search term for a stock with symbol match")
}
