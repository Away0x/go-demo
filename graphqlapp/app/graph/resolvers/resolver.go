package resolvers

import (
  "gorm.io/gorm"
  "graphqlapp/app/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
  ORM *gorm.DB
  UserServices services.IUserServices
}
