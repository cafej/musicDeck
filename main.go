package main

import (
	"musicDeck/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	//db := database.InitiateConnection()
	//defer db.Close()

	router := getEngine()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEngine() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handler.PingGet)
	router.GET("/songs", handler.GetSongs)

	return router
}
