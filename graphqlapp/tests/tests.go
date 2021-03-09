package tests

import (
  "graphqlapp/bootstrap"
  "os"
)

func SetupConfig() error {
  err := os.Chdir("../..")
  if err != nil {
    return err
  }
  bootstrap.SetupConfig("config/test.yaml", "yaml")
  return nil
}

func SetupDB() {
  bootstrap.SetupDB()
}

func SetupServer() {
  bootstrap.SetupServer()
}
