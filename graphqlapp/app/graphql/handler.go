package graphql

import (
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "graphqlapp/app/graphql/mutations"
  "graphqlapp/app/graphql/queries"
  "graphqlapp/app/services"
  "graphqlapp/core/context"
)

func NewHandler() context.AppHandlerFunc {
  us := services.NewUserServices()

  queryFields := mapFields(
    queries.NewUser(us),
  )

  mutationFields := mapFields(
    mutations.NewUser(us),
  )

  schema, err := graphql.NewSchema(
    graphql.SchemaConfig{
      Query: graphql.NewObject(graphql.ObjectConfig{
        Name: "Query",
        Fields: queryFields,
      }),
      Mutation: graphql.NewObject(graphql.ObjectConfig{
       Name: "Mutation",
       Fields: mutationFields,
      }),
    },
  )

  if err != nil {
    panic(err)
  }

  h := handler.New(&handler.Config{
    Schema:   &schema,
    Pretty:   true,
    GraphiQL: true,
  })

  return func(c *context.AppContext) error {
    h.ServeHTTP(c.Response().Writer, c.Request())
    return nil
  }
}

func mapFields(qs ...graphql.Fields) graphql.Fields {
  fields := make(graphql.Fields)
  all := make([]graphql.Fields, 0)
  all = append(all, qs...)

  for _, q := range all {
    for k, f := range q {
      fields[k] = f
    }
  }

  return fields
}
