package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	"goj/model"
)

func Seed(db *mongo.Database) {
	data, err := os.ReadFile("./migration/seed.json")
	if err != nil {
		panic(err)
	}
	// var albums model.Albums
	var albums []*model.Album
	err = json.Unmarshal(data, &albums)
	if err != nil {
		panic(err)
	}

	coll := db.Collection("albums")

	for _, album := range albums {
		res, err := coll.InsertOne(context.TODO(), album)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}

	// var manyModel []interface{}

	// for _, album := range albums {
	// 	manyModel = append(manyModel, album)
	// }

	// _, err = coll.InsertMany(context.TODO(), manyModel)
	// if err != nil {
	// 	panic(err)
	// }
}
