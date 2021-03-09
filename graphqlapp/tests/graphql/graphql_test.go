package graphql_test

import (
  "github.com/99designs/gqlgen/client"
  "graphqlapp/routes"
  "graphqlapp/tests"
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
  err = tests.SetupConfig()
  tests.SetupDB()

  graphqlClient = client.New(routes.NewGraphqlHandler())
  return
}

func after() {}

