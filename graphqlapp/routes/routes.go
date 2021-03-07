package routes

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "graphqlapp/core"
  "net/http"
  "strings"
)

func Register(router *core.Application) {
  staticURL := core.GetConfig().String("APP.STATIC_URL")
  publicDir := core.GetConfig().String("APP.PUBLIC_DIR")
  faviconURl := "/favicon.ico"

  if !core.GetConfig().IsDev() {
    router.Use(middleware.Recover())
  }

  if core.GetConfig().IsDev() {
    router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
      Format: "${status}   ${method}   ${latency_human}               ${uri}\n",
      Skipper: func(c echo.Context) bool {
        return strings.HasPrefix(c.Request().URL.Path, staticURL) || strings.HasPrefix(c.Request().URL.Path, faviconURl)
      },
    }))
  }

  router.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
    Getter: middleware.MethodFromForm("_method"),
  }))

  router.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
    RedirectCode: http.StatusMovedPermanently,
  }))

  if core.GetConfig().Bool("APP.GZIP") {
    router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
      Skipper: func(c echo.Context) bool {
        return !strings.HasPrefix(c.Request().URL.Path, staticURL)
      },
    }))
  }

  router.Static(staticURL, publicDir)
  router.File(faviconURl, publicDir+faviconURl)

  registerError(router)
  registerWeb(router)
  registerAPI(router)
  registerGraphql(router)
}
