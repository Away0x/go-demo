package config

import (
  "time"

  "github.com/spf13/viper"
)

const (
  defaultTempDir = "storage"
  defaultAppPort = "9999"
  defaultAppName = "server"
)

var defaultConfigMap = map[string]interface{}{
  // app
  "APP.NAME":          defaultAppName,
  "APP.VERSION":       "1.0.0",
  "APP.MODE":          "production", // dev mode
  "APP.ADDR":          ":" + defaultAppPort,
  "APP.URL":           "http://localhost:" + defaultAppPort,
  "APP.KEY":           "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
  "APP.TEMP_DIR":      defaultTempDir,
  "APP.PUBLIC_DIR":    "public",
  "APP.STATIC_URL":    "/public",
  "APP.UPLOAD_DIR":    "public/uploads",
  "APP.RESOURCES_DIR": "resources",
  "APP.TEMPLATE_DIR":  "resources/views",
  "APP.GZIP":          true,
  "APP.TEMPLATE_EXT":  "html",

  // db
  "DB.DEFAULT.CONNECTION":           "mysql",
  "DB.DEFAULT.HOST":                 "127.0.0.1",
  "DB.DEFAULT.PORT":                 "3306",
  "DB.DEFAULT.DATABASE":             defaultAppName,
  "DB.DEFAULT.USERNAME":             "root",
  "DB.DEFAULT.PASSWORD":             "",
  "DB.DEFAULT.OPTIONS":              "charset=utf8&parseTime=True&loc=Local",
  "DB.DEFAULT.MAX_OPEN_CONNECTIONS": 100,
  "DB.DEFAULT.MAX_IDLE_CONNECTIONS": 20,
  "DB.DEFAULT.AUTO_MIGRATE":         false,

  // jwt token
  "TOKEN.ACCESS_TOKEN_LIFETIME":  60 * time.Minute,
  "TOKEN.REFRESH_TOKEN_LIFETIME": 60 * 24 * time.Minute,

  // log
  "LOG.PREFIX":       "[APP]",
  "LOG.FOLDER":       defaultTempDir + "/logs/app",
  "LOG.LEVEL":        "debug", // log level: debug, info, warn, error, dpanic, panic, fatal
  "LOG.MAXSIZE":      10,
  "LOG.MAX_BACK_UPS": 5,
  "LOG.MAX_AGES":     30,

  // test
  "TEST.ENABLE_LOG": true,
}

func setupDefaultConfig() {
  for k, v := range defaultConfigMap {
    viper.SetDefault(k, v)
  }
}
