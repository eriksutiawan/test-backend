package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id primitive.ObjectID) (*User, error)
	FindByUserName(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user User) error
}
