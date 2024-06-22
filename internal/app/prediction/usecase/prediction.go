package usecase

import "context"

type Predictor interface {
	Predict(ctx context.Context, dto PredictDto) (*PredictResponse, error)
}
