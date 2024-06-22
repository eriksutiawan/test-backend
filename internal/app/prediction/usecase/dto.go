package usecase

import "mime/multipart"

type PredictDto struct {
	File   *multipart.FileHeader
	UserId string
}

type PredictResponse struct {
	UserId      string `json:"userId"`
	Predictions any    `json:"label"`
}
