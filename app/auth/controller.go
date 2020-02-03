package auth

import "github.com/gin-gonic/gin"

type Controller struct{}

func (a Controller) SignUp(c *gin.Context) {}

func (a Controller) SignIn(c *gin.Context) {}

func (a Controller) SignOut(c *gin.Context) {}
