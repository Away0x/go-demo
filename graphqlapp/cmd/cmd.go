package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "graphqlapp/bootstrap"
  "os"
)

const (
  defaultConfigFilePath = "config/development.yaml"
  configFileType = "yaml"
)

var configFilePath string

var rootCmd = &cobra.Command{
  Use: "graphqlapp",
}

// Execute execute cmd
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().
    StringVarP(&configFilePath, "config", "c", defaultConfigFilePath, "config file")
}

func initConfig() {
  if configFilePath == "" {
    configFilePath = defaultConfigFilePath
  }

  bootstrap.SetupConfig(configFilePath, configFileType)
}
