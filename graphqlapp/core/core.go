package core

import (
  "database/sql"
  "go.uber.org/zap"
  "gorm.io/gorm"
)

var (
  application       *Application
  defaultConnection *GormConnection
  appConfig         *AppConfig
  appLog            *zap.SugaredLogger
)

func GetApplication() *Application {
  if application == nil {
    panic("application is not initialized")
  }

  return application
}

func GetDefaultConnection() *sql.DB {
  if defaultConnection == nil || defaultConnection.DB == nil {
    panic("default connection is not initialized")
  }

  return defaultConnection.DB
}

func GetDefaultConnectionEngine() *gorm.DB {
  if defaultConnection == nil || defaultConnection.Engine == nil {
    panic("default connection is not initialized")
  }

  return defaultConnection.Engine
}

func GetConfig() *AppConfig {
  if appConfig == nil {
    panic("config is not initialized")
  }

  return appConfig
}

func GetLog() *zap.SugaredLogger {
  if appLog == nil {
    panic("log is not initialized")
  }

  return appLog
}
