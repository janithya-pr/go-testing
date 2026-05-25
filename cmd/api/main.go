package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// r.Use(
	// 	middleware.RequestID(),
	// 	middleware.Logging(log),
	// 	middleware.Recovery(),
	// )

	r.GET("/health", func(c *gin.Context) {

		// log.Info("health check")

		c.JSON(http.StatusOK, gin.H{
			"message": "healthy..",
		})
	})
	r.Run(":8080")
}