package usecase

import (
	"context"
)

type IUserGetter interface {
	GetUserByID(ctx context.Context, id string) (*UserResponse, error)
	GetUSerByUserName(ctx context.Context, userName string) (*UserResponse, error)
}

type IUserCreator interface {
	Create(ctx context.Context, user UserDto) error
}
