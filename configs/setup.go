package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(GetEnv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

var DB = ConnectDB()

func GetCollection(coll string) *mongo.Collection {
	collection := DB.Database("goserver").Collection(coll)
	return collection
}
