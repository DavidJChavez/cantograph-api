package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func main() {
	uri := "mongodb+srv://enimur:0JXqWyz7DIoXA6n4@cluster0.ibhswgv.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("Cluster0")
	albumsColl := db.Collection("albums")
	router := gin.Default()
	router.GET("/albums", func(ctx *gin.Context) {
		find, err := albumsColl.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}
		var result bson.M
		err = find.Decode(&result)
		if err == mongo.ErrNoDocuments {
			log.Println("No document was found")
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		if err != nil {
			panic(err)
		}

	})
	log.Println("Listen an serve on browser")
	router.Run(":8000")
}
