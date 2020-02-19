package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (a Controller) SignUp(c *gin.Context) {
	var signUpForm SignUpForm
	if err := c.ShouldBindJSON(&signUpForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
	return
}

func (a Controller) SignIn(c *gin.Context) {}

func (a Controller) SignOut(c *gin.Context) {}
