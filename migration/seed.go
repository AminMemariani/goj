package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"goj/model"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func Seed(client *mongo.Client) {

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

	coll := client.Database("gojdb").Collection("albums")

	res, err := coll.

	for _, album := range albums {
		res, err := coll.InsertOne(context.TODO(), album)
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
	}

}
