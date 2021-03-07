package cmd

import (
  "github.com/spf13/cobra"
  "graphqlapp/bootstrap"
)

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "run app server",
  Run: func(cmd *cobra.Command, args []string) {
    // init db
    bootstrap.SetupDB()
    // init server
    bootstrap.SetupServer()
    bootstrap.RunServer()
  },
}

func init() {
  rootCmd.AddCommand(serverCmd)
}
