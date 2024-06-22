package usecase

import "context"

type IAuth interface {
	Register(ctx context.Context, dto AuthRegisterDto) error
	Login(ctx context.Context, dto LoginDto) (*UserResponse, error)
}
