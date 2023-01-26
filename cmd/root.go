/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vijaynallagatla/cli-stocks/yapi"
)

var (
	apiClient *yapi.APIClient
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-stocks",
	Short: "Command line interface tool for fetching the Stocks summary",
	Long: `This command line interface uses Yahoo finance APIs and those are not officially
supported by Yahoo.
WARNING: These non-official APIs cannot be assumed stable and might break any time.
Also, you might violate Yahoo's terms of service. So use them at your own risk.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Print("error", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cfg := yapi.NewConfiguration()
	apiClient = yapi.NewAPIClient(cfg)
}
