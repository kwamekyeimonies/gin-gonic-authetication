package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB_Connection() *mongo.Client {

	Mongo_URL := os.Getenv("MONGO_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

	if err != nil{
		log.Fatal(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(), 100*time.Second)	
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Database Connected to Mongo DB")

	return client
}

var Client *mongo.Client = DB_Connection()


func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

	return collection
}
