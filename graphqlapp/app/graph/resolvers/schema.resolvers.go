package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphqlapp/app/graph/generated"
	gqlmodels "graphqlapp/app/graph/models"
	"graphqlapp/app/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *gqlmodels.CreateUser) (*models.User, error) {
	u := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	err := models.CreateModel(u)
	return u, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input *gqlmodels.UpdateUser) (*models.User, error) {
	u := new(models.User)
	err := models.DB().First(u, id).Error
	if err != nil {
		return nil, err
	}

	if input.Name != nil {
		u.Name = *input.Name
	}
	if input.Email != nil {
		u.Email = *input.Email
	}
	err = models.UpdateModel(u)
	return u, err
}

func (r *queryResolver) User(ctx context.Context, id *int) (*models.User, error) {
	u := new(models.User)
	err := models.DB().First(u, id).Error
	return u, err
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
