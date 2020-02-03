package server

import (
	auth_routes "gin-gorm-boilerplate/app/auth"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		auth_routes.NewRouter(v1)
	}
	return r
}
