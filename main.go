package main

import (
	"github.com/DavidJChavez/cantograph-api/database"
	"github.com/DavidJChavez/cantograph-api/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.DbConnection()
	router := gin.Default()
	router.GET("/albums", routes.GetAlbums)
	if err := router.Run(":8000"); err != nil {
		log.Panic(err)
	}
}
