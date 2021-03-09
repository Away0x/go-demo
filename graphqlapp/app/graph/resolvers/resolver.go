package resolvers

import (
  "context"
  "fmt"
  "gorm.io/gorm"
  "graphqlapp/app/services"
  appContext "graphqlapp/core/context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

const appContextKey = "__AppContextKey__"

type Resolver struct{
  ORM *gorm.DB
  UserServices services.IUserServices
}

func SaveAppContextToContext(c *appContext.AppContext) {
  ctx := context.WithValue(c.Request().Context(), appContextKey, c)
  c.SetRequest(c.Request().WithContext(ctx))
}

func (r *Resolver) AppContext(ctx context.Context) (*appContext.AppContext, error) {
  ac := ctx.Value(appContextKey)
  if ac == nil {
    return nil, fmt.Errorf("could not retrieve AppContext")
  }

  c, ok := ac.(*appContext.AppContext)
  if !ok {
    return nil, fmt.Errorf("AppContext has wrong type")
  }

  return c, nil
}

func (r *Resolver) MustAppContext(ctx context.Context) *appContext.AppContext {
  c, err := r.AppContext(ctx)
  if err != nil {
    panic(err)
  }
  return c
}
