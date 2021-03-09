package routes

import (
  "context"
  "github.com/99designs/gqlgen/graphql"
  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/99designs/gqlgen/graphql/playground"
  "github.com/vektah/gqlparser/v2/gqlerror"
  "graphqlapp/app/graph/generated"
  "graphqlapp/app/graph/resolvers"
  "graphqlapp/app/services"
  "graphqlapp/core"
  appContext "graphqlapp/core/context"
  "graphqlapp/core/errno"
)

const (
  graphqlPath = "/graphql"
  graphUIPath = "/graphql-ui"
)

func NewGraphqlHandler() *handler.Server {
  us := services.NewUserServices()

  return handler.NewDefaultServer(
    generated.NewExecutableSchema(
      generated.Config{
        Resolvers: &resolvers.Resolver{
          ORM: core.GetDefaultConnectionEngine(),
          UserServices: us,
        },
      },
    ),
  )
}

func registerGraphql(router *core.Application) {
  playgroundHandler := playground.Handler("GraphQL playground", graphqlPath)
  graphqlHandler := NewGraphqlHandler()

  graphqlHandler.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
    err := graphql.DefaultErrorPresenter(ctx, e)

    switch typed := err.Unwrap().(type) {
    case *errno.Errno:
      if err.Extensions == nil {
        err.Extensions = make(map[string]interface{})
      }

      err.Extensions["_code"] = typed.Code
      if typed.Data != nil {
        err.Extensions["_data"] = typed.Data
      }
    }

    return err
  })

  if core.GetConfig().IsDev() {
    router.RegisterHandler(router.GET, graphUIPath, func(c *appContext.AppContext) error {
      playgroundHandler.ServeHTTP(c.Response(), c.Request())
      return nil
    })
  }

  router.RegisterHandler(router.POST, graphqlPath, func(c *appContext.AppContext) error {
    resolvers.SaveAppContextToContext(c)
    graphqlHandler.ServeHTTP(c.Response(), c.Request())
    return nil
  })
}
