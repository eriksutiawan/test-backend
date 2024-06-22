package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Prediction struct {
	id        primitive.ObjectID `bson:"_id,omitempty"`
	userId    string             `bson:"userId"`
	imagePath string             `bson:"imagePath"`
	label     any                `bson:"label"`
}

func NewPrediction() *Prediction {
	return &Prediction{}
}

func (u *Prediction) SetImagePath(imagePath string) *Prediction {
	u.imagePath = imagePath
	return u
}

func (u *Prediction) SetLabel(label any) *Prediction {
	u.label = label
	return u
}

func (u *Prediction) SetUserId(userId string) *Prediction {
	u.userId = userId
	return u
}
