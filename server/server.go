package server

import (
	"goazl/config"
	"log"
	"strings"

	"goazl/server/handlers"

	"github.com/gin-gonic/gin"
)

func Run(c *config.Config) {
	port := c.Port

	if c.Environment != "" {
		if strings.ToLower(c.Environment) == "release" {
			log.Printf("Using environment: %v\n", c.Environment)
			gin.SetMode(gin.ReleaseMode)
		} else {
			gin.SetMode(gin.DebugMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	log.Printf("Server started on port: %v\n", port)

	router.GET("/lyrics", handlers.FetchLyrics())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": "PAGE_NOT_FOUND", "message": "Page not found",
		})
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server! - %v", err)
	}

	log.Println("Server running!")

}
