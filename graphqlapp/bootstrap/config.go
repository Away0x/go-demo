package bootstrap

import (
  "graphqlapp/config"
  "graphqlapp/core"
  "log"
)

func SetupConfig(configFilePath, configFileType string) {
  config.Setup(configFilePath, configFileType)
  core.NewAppConfig()

  if !core.GetConfig().IsTest() {
    err := config.WriteConfig(core.GetConfig().String("APP.TEMP_DIR") + "/config.json")
    if err != nil {
      log.Fatal(err)
    }
  }
}
