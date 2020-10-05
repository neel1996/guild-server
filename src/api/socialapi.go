package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/neel1996/guild-server/src/config"
	"github.com/neel1996/guild-server/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func SocialAPI(w http.ResponseWriter, r *http.Request) {
	var socialDataModel []model.SocialData
	socialDataModel = socialMongoModel()
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(socialDataModel)
}

func socialMongoModel() []model.SocialData {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongoData()))

	if err != nil {
		fmt.Println("Panic at mongo connect!")
		panic(err)
	}

	collection := client.Database(config.GetDBName()).Collection("social")
	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println("Panic at mongo Find!")
		panic(err)
	}

	var socialData []model.SocialData
	res.All(ctx, &socialData)

	fmt.Println("Data : ", socialData)

	return socialData
}
