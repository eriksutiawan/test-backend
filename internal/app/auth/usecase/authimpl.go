package usecase

import (
	"context"
	"errors"
	"os"
	useruscase "test-backend/internal/app/user/usecase"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	userCreator useruscase.IUserCreator
	userGetter  useruscase.IUserGetter
}

func NewAuth(userCreator useruscase.IUserCreator) *Auth {
	return &Auth{
		userCreator: userCreator,
	}
}

func (s Auth) Register(ctx context.Context, dto AuthRegisterDto) error {
	err := s.userCreator.Create(ctx, useruscase.UserDto{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s Auth) Login(ctx context.Context, dto LoginDto) (*UserResponse, error) {
	user, err := s.userGetter.GetUSerByUserName(ctx, dto.Username)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("password or username is wrong")
	}

	expired := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     expired,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Token:   tokenString,
		Expired: expired,
	}, nil
}
