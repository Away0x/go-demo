package config

import (
  "fmt"
  "github.com/spf13/viper"
  "strings"
)

type (
  Mode string
)

const (
  ModeProduction Mode = "production"
  ModeStaging Mode = "staging"
  ModeDevelopment Mode = "development"
  ModeTest Mode = "test"
)

// Setup Initialization configuration
// configFilePath: Configuration file path
// configFileType: Configuration file type
func Setup(configFilePath, configFileType string) {
  viper.SetConfigFile(configFilePath)
  viper.SetConfigType(configFileType)

  if err := viper.ReadInConfig(); err != nil {
    panic(
      fmt.Sprintf("Failed to read the configuration file. Please check if the %s configuration file exists: %v",
      configFilePath, err),
    )
  }

  // Setting configuration defaults
  setupDefaultConfig()

  // Setting environment variables: export APPNAME_APP_MODE=development)
  viper.AutomaticEnv()
  viper.SetEnvPrefix(viper.GetString("APP.NAME")) // 环境变量前缀
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

  fmt.Printf("\nConfiguration file loaded successfully: %s\n", configFilePath)
}

func WriteConfig(filename string) error {
  return viper.WriteConfigAs(filename)
}
