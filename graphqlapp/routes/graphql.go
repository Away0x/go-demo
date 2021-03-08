package routes

import (
  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/99designs/gqlgen/graphql/playground"
  "graphqlapp/app/graph/generated"
  "graphqlapp/app/graph/resolvers"
  "graphqlapp/app/services"
  "graphqlapp/core"
  "graphqlapp/core/context"
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

  if core.GetConfig().IsDev() {
    router.RegisterHandler(router.GET, graphUIPath, func(c *context.AppContext) error {
      playgroundHandler.ServeHTTP(c.Response(), c.Request())
      return nil
    })
  }

  router.RegisterHandler(router.POST, graphqlPath, func(c *context.AppContext) error {
    graphqlHandler.ServeHTTP(c.Response(), c.Request())
    return nil
  })
}
