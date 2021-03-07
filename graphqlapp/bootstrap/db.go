package bootstrap

import (
  "graphqlapp/core"
  "graphqlapp/database"
  "log"
)

func SetupDB() {
  db, sqlDB := database.SetupDefaultDatabase()

  // 自动迁移
  if core.GetConfig().Bool("DB.DEFAULT.AUTO_MIGRATE") {
    err := db.AutoMigrate(database.RegisterAutoMigrateModel()...)
    if err != nil {
      log.Fatal(err)
    }
  }

  core.NewDefaultConnection(db, sqlDB)
}
