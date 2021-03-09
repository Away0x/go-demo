package database

import (
  "database/sql"
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "graphqlapp/app/models"
  "graphqlapp/core"
  "graphqlapp/core/pkg/dbutils"
)

// SetupDefaultDatabase 初始化默认的数据库
func SetupDefaultDatabase() (*gorm.DB, *sql.DB) {
  dsn := dbutils.BuildDatabaseDSN(core.GetConfig().DefaultString("DB.DEFAULT.CONNECTION", "mysql"), dbutils.DatabaseConfig{
    UserName: core.GetConfig().String("DB.DEFAULT.USERNAME"),
    Password: core.GetConfig().String("DB.DEFAULT.PASSWORD"),
    Host:     core.GetConfig().String("DB.DEFAULT.HOST"),
    Port:     core.GetConfig().String("DB.DEFAULT.PORT"),
    DBName:   core.GetConfig().String("DB.DEFAULT.DATABASE"),
    Options:  core.GetConfig().String("DB.DEFAULT.OPTIONS"),
  }, func(s string) string {
    return core.GetConfig().String("DB.DEFAULT.DATABASE")
  })

  dd, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(getGormLoggerLevel()),
  })
  if err != nil {
    panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
  }

  sqlDB, err := dd.DB()
  if err != nil {
    panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
  }

  sqlDB.SetMaxOpenConns(core.GetConfig().Int("DB.DEFAULT.MAX_OPEN_CONNECTIONS"))
  sqlDB.SetMaxIdleConns(core.GetConfig().Int("DB.DEFAULT.MAX_IDLE_CONNECTIONS"))

  fmt.Printf("\n(SETUP) Default database connection successful: %s\n", dsn)
  return dd, sqlDB
}

func RegisterAutoMigrateModel() []interface{} {
  return []interface{}{
    &models.User{},
  }
}

func getGormLoggerLevel() logger.LogLevel {
  if core.GetConfig().IsDev() {
    return logger.Info
  }

  return logger.Silent
}
