package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PredictionRepository struct {
	collection *mongo.Collection
}

func NewPredictionRepository(database *mongo.Database) *PredictionRepository {
	collection := database.Collection("prediction")
	return &PredictionRepository{
		collection: collection,
	}
}

func (r *PredictionRepository) Creator(ctx context.Context, model Prediction) (*Prediction, error) {
	model.id = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, model)
	return &model, err
}
