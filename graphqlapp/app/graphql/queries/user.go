package queries

import (
  "github.com/graphql-go/graphql"
  "graphqlapp/app/graphql/types"
  "graphqlapp/app/services"
)

func NewUser(us services.IUserServices) graphql.Fields {
  return graphql.Fields{
    "user": &graphql.Field{
      Name:        "user detail",
      Description: "user detail",
      Type:        types.User,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.Int,
        },
      },
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return us.Detail(uint(p.Args["id"].(int)))
      },
    },
    "userList": &graphql.Field{
      Name:        "user list",
      Description: "user list",
      Type:        graphql.NewList(types.User),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return us.List()
      },
    },
  }
}
