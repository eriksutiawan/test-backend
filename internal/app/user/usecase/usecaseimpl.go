package usecase

import (
	"context"
	"test-backend/internal/app/user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.IUserRepository
}

func NewUser(repo repository.IUserRepository) *User {
	return &User{
		repo: repo,
	}
}

func (s User) Create(ctx context.Context, user UserDto) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userData := *repository.NewUser().
		SetEmail(user.Email).
		SetUsername(user.Email).
		SetPassword(string(hashedPassword))

	s.repo.Create(ctx, userData)
	return nil
}

func (s User) GetUserByID(ctx context.Context, id string) (*UserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Id:       user.GetId().Hex(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}, nil
}

func (s User) GetUSerByUserName(ctx context.Context, userName string) (*UserResponse, error) {
	user, err := s.repo.FindByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Id:       user.GetId().Hex(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}, nil
}
