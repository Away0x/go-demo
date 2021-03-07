package core

import (
  "encoding/json"
  "github.com/labstack/echo/v4"
  "graphqlapp/core/context"
  "io/ioutil"
  "strings"
)

type Application struct {
  *echo.Echo
}

func NewApplication(e *echo.Echo) {
  application = &Application{
    Echo: e,
  }
}

func (a *Application) RoutePath(name string, params ...interface{}) string {
  return a.Reverse(name, params...)
}

func (a *Application) PrintRoutes(filename string) error {
  routes := make([]*echo.Route, 0)
  for _, item := range a.Routes() {
    if strings.HasPrefix(item.Name, "github.com") ||
      strings.HasSuffix(item.Name, "notFoundHandler") {
      continue
    }

    routes = append(routes, item)
  }

  routesStr, _ := json.MarshalIndent(struct {
    Count  int           `json:"count"`
    Routes []*echo.Route `json:"routes"`
  }{
    Count:  len(routes),
    Routes: routes,
  }, "", "  ")

  return ioutil.WriteFile(filename, routesStr, 0644)
}

func (a *Application) RegisterRoutes(register func(*Application)) {
  a.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      cc := &context.AppContext{Context: c}
      return hf(cc)
    }
  })

  register(a)
}

func (a *Application) RegisterHandler(
  fn context.EchoRegisterFunc,
  path string, h context.AppHandlerFunc,
  m ...echo.MiddlewareFunc) *echo.Route {
  if path != "" && !strings.HasPrefix(path, "/") {
    path = "/" + path
  }

  return fn(path, func(c echo.Context) error {
    cc, ok := c.(*context.AppContext)
    if !ok {
      cc = context.NewAppContext(c)
      return h(cc)
    }
    return h(cc)
  }, m...)
}
