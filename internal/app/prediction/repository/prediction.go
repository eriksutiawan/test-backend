package repository

import "context"

type IRepository interface {
	Creator(ctx context.Context, model Prediction) (*Prediction, error)
}
