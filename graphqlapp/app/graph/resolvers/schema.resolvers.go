package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphqlapp/app/graph/generated"
	"graphqlapp/app/graph/gqlmodels"
	"graphqlapp/app/models"
	"graphqlapp/core/errno"
)

func (r *mutationResolver) Login(ctx context.Context, input gqlmodels.DoLogin) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *gqlmodels.CreateUser) (*models.User, error) {
	u := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	err := models.CreateModel(u)
	return u, errno.DatabaseErr.WithErr(err)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input *gqlmodels.UpdateUser) (*models.User, error) {
	u := new(models.User)
	err := models.DB().First(u, id).Error
	if err != nil {
		return nil, err
	}

	u.Name = *input.Name
	u.Email = *input.Email
	err = models.UpdateModel(u)
	return u, errno.DatabaseErr.WithErr(err)
}

func (r *queryResolver) User(ctx context.Context, id *int) (*models.User, error) {
	u := new(models.User)
	err := models.DB().First(u, id).Error
	return u, errno.NotFoundErr.WithErr(err)
}

func (r *queryResolver) Users(ctx context.Context, page *int, perPage *int) (*gqlmodels.UserList, error) {
	items, total, err := r.UserServices.List(*page, *perPage)
	if err != nil {
		return nil, errno.NotFoundErr.WithErr(err)
	}

	return &gqlmodels.UserList{
		Items:    items,
		PageInfo: &gqlmodels.PageResult{Total: int(total), Page: *page, PerPage: *perPage},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
