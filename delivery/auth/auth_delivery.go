package auth_delivery

import "github.com/gin-gonic/gin"

type AuthDelivery interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
