package routes

import (
  "graphqlapp/app/graphql"
  "graphqlapp/core"
)

func registerGraphql(router *core.Application) {
  graphqlHandler := graphql.NewHandler()

  if core.GetConfig().IsDev() {
    router.RegisterHandler(router.GET, "/graphql", graphqlHandler)
  }

  router.RegisterHandler(router.POST, "/graphql", graphqlHandler)
}
