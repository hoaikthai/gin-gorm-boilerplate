package middlewares

import (
	"gin-gorm-boilerplate/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		request_api_key := c.Request.Header.Get("Api-Key")
		config := config.GetConfig()
		if request_api_key != config.GetString("api_key") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  401,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
