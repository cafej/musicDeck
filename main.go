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
	//router.GET("/artist/:artist", handler.GetSongsByArtist)
	router.GET("/song/:query", handler.GetSongsBy("song"))
	router.GET("/artist/:query", handler.GetSongsBy("artist"))
	router.GET("/genre/:query", handler.GetSongsByGenre)
	router.GET("length/:minimum/:maximum", handler.GetSongsByLength)
	router.GET("/total", handler.GetGenresTime)
	return router
}
