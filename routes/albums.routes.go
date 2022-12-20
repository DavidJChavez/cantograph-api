package routes

import (
	"context"
	"github.com/DavidJChavez/cantograph-api/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetAlbums(ctx *gin.Context) {
	albumsCollection := database.DB.Collection("albums")
	cur, err := albumsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := cur.Close(ctx); err != nil {
			log.Panic(err)
		}
	}()
	var result []bson.M
	for cur.Next(context.TODO()) {
		var album bson.M
		err := cur.Decode(&album)
		if err != nil {
			log.Panic(err)
		}
		result = append(result, album)
	}
	if err := cur.Err(); err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, result)
}
