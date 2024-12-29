package auth_delivery

import (
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

// CreateTags		godoc
// @Summary			Logout endpoint.
// @Tags				auth
// @Success			200 {object} dto.ResLogout "Ok"
// @Router			/auth/logout [post]
func (delivery *AuthDeliveryImpl) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "/", "", false, true)
	c.JSON(http.StatusOK, dto.ResLogout{Message: "Logged out successfully"})
}
