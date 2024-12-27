package auth_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

func (delivery *AuthDeliveryImpl) Login(c *gin.Context) {
	var login dto.ReqLogin

	// Parse json request
	if err := c.ShouldBindJSON(&login); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Login (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	httpCode, user, jwtToken, err := delivery.AuthUsecase.Login(c.Request.Context(), login)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Login (%s)", msg)
		c.String(httpCode, msg)
		return
	}

	c.SetCookie("jwt", jwtToken, (15 * 24 * 60 * 60), "/", "", false, true)
	c.JSON(http.StatusOK, user)
}
