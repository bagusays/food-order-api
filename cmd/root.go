package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "food-order-api",
	Short: "food-order-api is xxx xxx xxx",
}

func Execute() error {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(migrateCmd)

	return rootCmd.Execute()
}
