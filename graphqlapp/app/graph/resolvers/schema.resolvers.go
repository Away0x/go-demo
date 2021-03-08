package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphqlapp/app/graph/generated"
	gqlmodels "graphqlapp/app/graph/models"
	"graphqlapp/app/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *gqlmodels.NewUser) (*models.User, error) {
	return r.UserServices.Create(&models.User{
	  Name: input.Name,
	  Email: input.Email,
  })
}

func (r *queryResolver) User(ctx context.Context, id *int) (*models.User, error) {
	if id != nil {
	  return r.UserServices.Detail(*id)
  }

  return nil, fmt.Errorf("self user")
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserServices.List()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
