package routes

import (
  "graphqlapp/core"
  "graphqlapp/core/context"
  "time"
)

func registerWeb(router *core.Application) {
  router.RegisterHandler(router.GET, "welcome", func(c *context.AppContext) error {
    now := time.Now()

    return c.SuccessHTML("welcome", context.RespData{
      "time": now.Format("2006-01-02"),
    })
  }).Name = "welcome"
}
