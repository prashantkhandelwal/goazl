package handlers

import (
	"net/http"
	"strings"

	"goazl/server/azl"

	"github.com/gin-gonic/gin"
)

func FetchLyrics() gin.HandlerFunc {
	fn := func(g *gin.Context) {

		artistname := azl.Format(g.Query("artist"))
		songname := azl.Format(g.Query("song"))

		if len(strings.TrimSpace(artistname)) > 0 || len(strings.TrimSpace(songname)) > 0 {
			r := azl.FetchLyrics(artistname, songname)
			g.JSON(http.StatusOK, gin.H{"result": r})
		} else {
			g.JSON(http.StatusBadRequest, gin.H{"result": "Error: artist or song name cannot be blank."})
		}
	}
	return gin.HandlerFunc(fn)
}
