package bootstrap

import (
  "fmt"
  "github.com/flosch/pongo2"
  "github.com/labstack/echo/v4"
  "graphqlapp/core"
  "graphqlapp/core/pkg/strutils"
  "graphqlapp/core/pkg/tplutils"
  "graphqlapp/routes"
  "log"
)

func SetupServer() {
  // setup log
  core.SetupLog()

  e := echo.New()
  e.Debug = core.GetConfig().IsDev()
  e.HideBanner = true
  core.NewApplication(e)
  // register routes
  core.GetApplication().RegisterRoutes(routes.Register)
  err := core.GetApplication().PrintRoutes(core.GetConfig().String("APP.TEMP_DIR") + "/routes.json")
  if err != nil {
    log.Fatal(err)
  }

  SetupServerRender()

  fmt.Printf(
    "\n(SETUP) app(%s) mode is %s, %s\n\n",
    core.GetConfig().String("APP.NAME"),
    core.GetConfig().AppMode(),
    core.GetConfig().String("APP.URL"),
  )
}

func RunServer() {
  core.GetApplication().Echo.Logger.Fatal(core.GetApplication().Start(core.GetConfig().String("APP.ADDR")))
}

func SetupServerRender() {
  render := tplutils.NewRenderer()
  tplutils.SetupTpl(&tplutils.Config{
    GetRoutePath: core.GetApplication().RoutePath,
    GeneratePublicPath: func(path string) string {
      staticURL := core.GetConfig().String("APP.STATIC_URL")
      if core.GetConfig().IsDev() {
        return fmt.Sprintf("%s%s?v=%s", staticURL, path, strutils.RandomCreateBytes(6))
      }
      return fmt.Sprintf("%s%s", staticURL, path)
    },
  })

  // template dir
  render.AddDirectory(core.GetConfig().String("APP.TEMPLATE_DIR"))

  // template global var
  globalVar := pongo2.Context{
    "APP_NAME":    core.GetConfig().String("APP.NAME"),
    "APP_MODE": string(core.GetConfig().AppMode()),
    "APP_URL":     core.GetConfig().String("APP.URL"),
  }

  render.UseContextProcessor(func(c echo.Context, p pongo2.Context) {
    p.Update(globalVar)

    data := pongo2.Context{
      "route_path": c.Request().URL.Path,
    }

    p.Update(data)
  })

  core.GetApplication().Renderer = render

  // tags
  pongo2.RegisterTag("route", tplutils.RouteTag)
  pongo2.RegisterTag("static", tplutils.StaticTag)
}

