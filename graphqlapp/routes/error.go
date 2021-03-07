package routes

import (
  "github.com/labstack/echo/v4"
  "graphqlapp/core"
  "graphqlapp/core/context"
  "graphqlapp/core/errno"
  "log"
  "net/http"
)

func registerError(router *core.Application) {
  echo.NotFoundHandler = notFoundHandler
  echo.MethodNotAllowedHandler = notFoundHandler

  router.HTTPErrorHandler = func(err error, c echo.Context) {
    errnoData := transformErrorType(err)

    if !c.Response().Committed {
      if c.Request().Method == http.MethodHead {
        err = c.NoContent(http.StatusOK)
      } else {
        cc := context.NewAppContext(c)
        err = cc.ErrorJSON(errnoData)
      }
      if err != nil {
        log.Printf("routes/error#HTTPErrorHandler: %s", err)
      }
    }
  }
}

func transformErrorType(err error) *errno.Errno {
  switch typed := err.(type) {
  case *errno.Errno:
    return typed
  default:
    return errno.UnknownErr.WithErr(typed).(*errno.Errno)
  }
}

func notFoundHandler(c echo.Context) error {
  return errno.NotFoundErr
}
