package auth_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *AuthDeliveryImpl) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "/", "", false, true)
	c.String(http.StatusOK, "Logged out successfully")
}
