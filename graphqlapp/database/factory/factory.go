package factory

import (
  "fmt"
  "graphqlapp/core"
)

func dropAndCreateTable(table interface{}) {
  err := core.GetDefaultConnectionEngine().Migrator().DropTable(table)
  if err != nil {
    panic(err)
  }
  err = core.GetDefaultConnectionEngine().Migrator().CreateTable(table)
  if err != nil {
    panic(err)
  }
}

// Run run database factory
func Run() {
  usersTableSeeder()
  fmt.Println("database.factory running")
}
