package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Cresdentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS, POST, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Api-Key, Accept-Encoding, X-CSRF-Token, Cache-Control")

		c.Next()
	}
}
