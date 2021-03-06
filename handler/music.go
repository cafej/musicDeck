package handler

import (
	"musicDeck/database"
	"musicDeck/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSongs will retrive all songs in the DB
func GetSongs(c *gin.Context) {
	db := database.InitiateConnection()
	songs := []structs.Song{}
	rows, err := db.Query(`
	SELECT s.ID, s.artist, s.song, s.length, g.name, g.ID as genreID 
	FROM songs s INNER JOIN  genres g ON s.genre = g.ID;
	`)
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(400, err)
	}
	var song structs.Song
	for rows.Next() {
		// we could prettify this line with sqlx
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre.Name, &song.Genre.ID)
		if err != nil {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.AbortWithError(400, err)
		}
		songs = append(songs, song)
	}
	defer db.Close()
	c.JSON(200, songs)
}

// GetSongsByArtist will retrive all songs in the DB that matches the artist in the request
func GetSongsByArtist(c *gin.Context) {
	artistName := c.Param("artist")
	db := database.InitiateConnection()
	songs := []structs.Song{}
	query := `
	SELECT s.ID, s.artist, s.song, s.length, g.name, g.ID as genreID 
	FROM songs s INNER JOIN  genres g ON s.genre = g.ID WHERE s.artist LIKE ?;
	`
	rows, err := db.Query(query, "%"+artistName+"%")
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(400, err)
	}
	var song structs.Song
	for rows.Next() {
		// we could prettify this line with sqlx
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre.Name, &song.Genre.ID)
		if err != nil {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.AbortWithError(400, err)
		}
		songs = append(songs, song)
	}
	defer db.Close()
	c.JSON(200, songs)
}

// GetSongsBy will retrive all songs in the DB that matches the artist in the request and type
func GetSongsBy(queryType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchParam := c.Param("query")
		db := database.InitiateConnection()
		songs := []structs.Song{}
		sql := `
		SELECT s.ID, s.artist, s.song, s.length, g.name, g.ID as genreID 
		FROM songs s INNER JOIN  genres g ON s.genre = g.ID WHERE s.` + queryType + ` LIKE '%` + searchParam + `%';`
		rows, err := db.Query(sql)
		if err != nil {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.AbortWithError(400, err)
		}
		var song structs.Song
		for rows.Next() {
			// we could prettify this line with sqlx
			err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre.Name, &song.Genre.ID)
			if err != nil {
				c.Header("Content-Type", "application/json; charset=utf-8")
				c.AbortWithError(400, err)
			}
			songs = append(songs, song)
		}
		defer db.Close()
		c.JSON(200, songs)
	}
}

// GetSongsByGenre will retrive all songs in the DB that matches the genre in the request
func GetSongsByGenre(c *gin.Context) {
	genreName := c.Param("query")
	db := database.InitiateConnection()
	songs := []structs.Song{}
	query := `
	SELECT s.ID, s.artist, s.song, s.length, g.name, g.ID as genreID 
	FROM songs s INNER JOIN  genres g ON s.genre = g.ID WHERE g.name LIKE ?;
	`
	rows, err := db.Query(query, "%"+genreName+"%")
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(400, err)
	}
	var song structs.Song
	for rows.Next() {
		// we could prettify this line with sqlx
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre.Name, &song.Genre.ID)
		if err != nil {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.AbortWithError(400, err)
		}
		songs = append(songs, song)
	}
	defer db.Close()
	c.JSON(200, songs)
}

// GetSongsByLength will retrive all songs in the DB that are between the range provided
func GetSongsByLength(c *gin.Context) {
	min := c.Param("minimum")
	max := c.Param("maximum")
	db := database.InitiateConnection()
	songs := []structs.Song{}
	sql := `
	SELECT s.ID, s.artist, s.song, s.length, g.name, g.ID as genreID 
	FROM songs s INNER JOIN  genres g ON s.genre = g.ID WHERE s.length BETWEEN ` + min + ` and ` + max + `;`
	rows, err := db.Query(sql)
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(400, err)
	}
	var song structs.Song
	for rows.Next() {
		// we could prettify this line with sqlx
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre.Name, &song.Genre.ID)
		if err != nil {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.AbortWithError(400, err)
		}
		songs = append(songs, song)
	}
	defer db.Close()
	c.JSON(200, songs)

}

// GetGenresTime will retrive all genres in the DB with the total time of the songs summed
func GetGenresTime(c *gin.Context) {
	db := database.InitiateConnection()
	genres := []structs.GenreTotal{}
	rows, err := db.Query(`
	SELECT g.name, sum(s.length) as totalTime, count(s.song) as totalSongs
	FROM songs s INNER JOIN  genres g ON s.genre = g.ID 
	GROUP BY g.name;
	`)
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.AbortWithError(400, err)
	}
	for rows.Next() {
		genre := structs.GenreTotal{}
		err = rows.Scan(&genre.Name, &genre.TotalTime, &genre.TotalSongs)
		if err != nil {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.AbortWithError(400, err)
		}
		genres = append(genres, genre)
	}
	defer db.Close()
	c.JSON(200, genres)
}

// PingGet will be used to test that the API is working
func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}
