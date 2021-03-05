package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSongs will retrive all songs in the DB
func GetSongs(c *gin.Context) {

}

// PingGet will be used to test that the API is working
func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "found me",
	})
}
