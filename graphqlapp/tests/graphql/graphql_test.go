package graphql_test

import (
  "github.com/99designs/gqlgen/client"
  "graphqlapp/bootstrap"
  "graphqlapp/routes"
  "os"
  "testing"
)

var graphqlClient *client.Client

func TestMain(m *testing.M)  {
  if err := before(); err != nil {
    panic(err)
  }
  m.Run()
  after()
}

func before() (err error) {
  err = os.Chdir("../..")
  bootstrap.SetupConfig("config/test.yaml", "yaml")
  bootstrap.SetupDB()

  graphqlClient = client.New(routes.NewGraphqlHandler())
  return
}

func after() {}

