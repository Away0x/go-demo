package routes

import (
  "github.com/labstack/echo/v4/middleware"
  "graphqlapp/core"
  "graphqlapp/core/context"
)

const (
  APIPrefix = "/api"
)

func registerAPI(router *core.Application) {
  e := router.Group(APIPrefix, middleware.CORS())

  router.RegisterHandler(e.GET, "test", func(c *context.AppContext) error {
    core.GetLog().Debug("log test")
    return c.SuccessJSON(context.RespData{
      "hello": "world",
    })
  }).Name = "test"
}
