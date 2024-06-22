package usecase

import (
	"context"
	"encoding/json"
	"os/exec"
	"test-backend/internal/app/prediction/repository"
	"test-backend/internal/pkg/utils"
)

type Prediction struct {
	repo repository.IRepository
}

func NewPrediction(repo repository.IRepository) *Prediction {
	return &Prediction{
		repo: repo,
	}
}

func (s Prediction) Predict(ctx context.Context, dto PredictDto) (*PredictResponse, error) {
	filepath, err := utils.SaveImage(dto.File)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("python", "predict.py", filepath)
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var predictions any
	if err := json.Unmarshal(out, &predictions); err != nil {
		return nil, err
	}

	predict := *repository.NewPrediction().
		SetImagePath(filepath).
		SetLabel(predictions).
		SetUserId(dto.UserId)

	_, err = s.repo.Creator(ctx, predict)
	if err != nil {
		return nil, err
	}

	return &PredictResponse{
		Predictions: predictions,
		UserId:      dto.UserId,
	}, nil
}
