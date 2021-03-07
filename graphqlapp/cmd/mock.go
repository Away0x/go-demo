package cmd

import (
  "github.com/spf13/cobra"
  "graphqlapp/bootstrap"
  "graphqlapp/database/factory"
)

var mockCmd = &cobra.Command{
  Use:   "mock",
  Short: "mock data",
  Run: func(cmd *cobra.Command, args []string) {
    bootstrap.SetupDB()
    factory.Run()
  },
}

func init() {
  rootCmd.AddCommand(mockCmd)
}
