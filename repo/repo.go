package repo

import (
	"context"
	"goj/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAlbums(db *mongo.Database) (albums []*model.Album, err error) {
	var cursor *mongo.Cursor
	cursor, err = db.Collection("albums").Find(context.TODO(), bson.D{})
	if err != nil {
		return
	}

	err = cursor.All(context.TODO(), &albums)
	if err != nil {
		return
	}

	return
}
