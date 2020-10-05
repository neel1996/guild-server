package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/neel1996/guild-server/src/config"
	"github.com/neel1996/guild-server/src/model"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func PersonalDataApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var personalDataModel []model.PersonalData
	personalDataModel = personalMongoData()

	json.NewEncoder(w).Encode(personalDataModel[0])
}

func personalMongoData() []model.PersonalData {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongoData()))

	if err != nil {
		fmt.Println("Error in Mongo connect : ", err)
		panic(err)
	}

	collection := client.Database(config.GetDBName()).Collection("personal")
	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println("Error in Mongo Find", err)
		panic(err)
	}

	var personalData []model.PersonalData
	res.All(ctx, &personalData)

	fmt.Println("Data : ", personalData)

	return personalData
}
