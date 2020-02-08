package server

import (
	auth "gin-gorm-boilerplate/app/auth"
	"gin-gorm-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.ApiMiddleware())
	v1 := r.Group("/v1")
	{
		auth.NewRouter(v1)
	}
	return r
}
