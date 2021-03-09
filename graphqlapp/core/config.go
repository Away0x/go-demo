package core

import (
  "github.com/spf13/viper"
  "graphqlapp/config"
  "time"
)

type AppConfig struct{}

func NewAppConfig() {
  appConfig = &AppConfig{}
}

func (*AppConfig) String(key string) string {
  return viper.GetString(key)
}

func (*AppConfig) DefaultString(key string, defaultVal string) string {
  v := viper.GetString(key)
  if v == "" {
    return defaultVal
  }

  return v
}

func (*AppConfig) Int(key string) int {
  return viper.GetInt(key)
}

func (*AppConfig) DefaultInt(key string, defaultVal int) int {
  v := viper.GetInt(key)
  if v == 0 {
    return defaultVal
  }

  return v
}

func (*AppConfig) Bool(key string) bool {
  return viper.GetBool(key)
}

func (*AppConfig) Duration(key string) time.Duration {
  return viper.GetDuration(key)
}

func (c *AppConfig) IsDev() bool {
  return c.AppMode() == config.ModeDevelopment
}

func (c *AppConfig) IsTest() bool {
  return c.AppMode() == config.ModeTest
}

func (c *AppConfig) AppMode() config.Mode {
  mode := config.Mode(c.String("APP.Mode"))

  switch mode {
  case config.ModeProduction:
    return config.ModeProduction
  case config.ModeStaging:
    return config.ModeStaging
  case config.ModeDevelopment:
    return config.ModeDevelopment
  case config.ModeTest:
    return config.ModeTest
  default:
    return config.ModeDevelopment
  }
}
