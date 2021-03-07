package types

import "github.com/graphql-go/graphql"

var User = graphql.NewObject(
  graphql.ObjectConfig{
    Name: "User",
    Description: "Users data",
    Fields: graphql.Fields{
      "name":      &graphql.Field{Type: graphql.String},
      "email":     &graphql.Field{Type: graphql.String},
    },
  },
)
