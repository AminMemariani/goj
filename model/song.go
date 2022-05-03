package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty"`
	AudioUrl string             `bson:"audio_url,omitempty" json:"audio_url,omitempty"`
	ImageUrl string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
}
