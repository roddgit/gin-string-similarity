package configs

import (
	"context"
	"fmt"
	"gin-string-similarity/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var logs = ZeroLogger()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("api_db").Collection(collectionName)

	return collection
}

func InsertTresholdMongo(logs_id string, req_pmo_raw string, req_core_raw string, req_pmo string, req_core string, treshold float64) string {
	var eformMatchCollection *mongo.Collection = GetCollection(ConnectDB(), "compare_name")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newCompare := models.Eform_matching_treshold{
		Logs_Id:       logs_id,
		Name_PMO_Raw:  req_pmo_raw,
		Name_Core_Raw: req_core_raw,
		Name_PMO:      req_pmo,
		Name_Core:     req_core,
		Treshold:      treshold,
		Create_Date:   DateTimeNow(),
	}

	insert, err := eformMatchCollection.InsertOne(ctx, newCompare)

	if err != nil {
		logs.Error().Str("Fault", ArrayToString(err)).Msg(logs_id)
	}

	result := fmt.Sprint(insert)

	return result
}
