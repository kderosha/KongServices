package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"context"
)


// Instantiate a new mongo client 
// https://www.mongodb.com/docs/drivers/go/current
func NewDb() *mongo.Client {
	mongoUri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil{
		panic(err)
	}
	return client;
}