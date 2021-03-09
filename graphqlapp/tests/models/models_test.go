package models_test

import (
  "graphqlapp/tests"
  "testing"
)

func TestMain(m *testing.M)  {
  if err := before(); err != nil {
    panic(err)
  }
  m.Run()
  after()
}

func before() (err error) {
  err = tests.SetupConfig()
  tests.SetupDB()
  return
}

func after() {}
