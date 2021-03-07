package mutations

import (
  "github.com/graphql-go/graphql"
  "graphqlapp/app/graphql/types"
  "graphqlapp/app/models"
  "graphqlapp/app/services"
)

func NewUser(us services.IUserServices) graphql.Fields {
 return graphql.Fields{
   "createUser": &graphql.Field{
     Type: types.User,
     Description: "create user",
     Args: graphql.FieldConfigArgument{
       "name": &graphql.ArgumentConfig{Type: graphql.String},
       "email": &graphql.ArgumentConfig{Type: graphql.String},
     },
     Resolve: func(p graphql.ResolveParams) (interface{}, error) {
       u := &models.User{
         Name: p.Args["name"].(string),
         Email: p.Args["email"].(string),
       }

       return us.Create(u)
     },
   },
 }
}
