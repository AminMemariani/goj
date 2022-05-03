package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// type Albums struct {
// 	Albums []*Album `json:"albums,omitempty"`
// }

type Album struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string             `bson:"title" json:"title,omitempty"`
	Desc  string             `bson:"desc" json:"desc,omitempty"`
	Songs []*Song            `bson:"songs" json:"songs,omitempty"`
}
