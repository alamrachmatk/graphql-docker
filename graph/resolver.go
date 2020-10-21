package graph

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"dexp/graph/generated"
	"dexp/graph/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Email: input.Email,
	}

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var res []*model.User

	user1 := &model.User{
		ID: 1,
		FirstName: "Alam",
		LastName: "Rachmat",
		Email: "alam@gmai.com",
		CreatedAt: 0,
		UpdatedAt: 0,
	}

	user2 := &model.User{
		ID: 2,
		FirstName: "Alam 2",
		LastName: "Rachmat 2",
		Email: "alam@gmai.com",
		CreatedAt: 0,
		UpdatedAt: 0,
	}
	
	res = append(res, user1)
	res = append(res, user2)
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }