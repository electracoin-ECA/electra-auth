package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors sets the CORS headers.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
