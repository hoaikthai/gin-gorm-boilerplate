package middlewares

import (
	"gin-gorm-boilerplate/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		request_api_key := c.Request.Header.Get("Api-Key")
		config := config.GetConfig()
		if request_api_key != config.GetString("server.api_key") {
			log.Print("Unauthorized request from " + c.ClientIP())
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized request",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
