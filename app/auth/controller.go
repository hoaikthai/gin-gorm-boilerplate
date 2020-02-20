package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthService interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
}

type AuthController struct{}

func (a AuthController) SignUp(c *gin.Context) {
	var signUpForm SignUpForm
	if err := c.ShouldBindJSON(&signUpForm); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	if isExist(signUpForm) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ERRORS.AUTH.EXISTING_USER",
		})
		return
	}
	if err := createUser(signUpForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
	return
}

func (a AuthController) SignIn(c *gin.Context) {}

func (a AuthController) SignOut(c *gin.Context) {}
