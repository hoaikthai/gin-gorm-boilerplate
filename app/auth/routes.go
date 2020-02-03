package auth

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.RouterGroup) {
	controller := new(Controller)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", controller.SignUp)
		authGroup.POST("/signin", controller.SignIn)
		authGroup.POST("/signout", controller.SignOut)
	}
}
