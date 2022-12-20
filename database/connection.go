package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Database

func DbConnection() {
	uri := "mongodb+srv://enimur:0JXqWyz7DIoXA6n4@cluster0.ibhswgv.mongodb.net/?retryWrites=true&w=majority"
	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if e != nil {
		log.Panic(e)
	}
	DB = client.Database("develop")
}
