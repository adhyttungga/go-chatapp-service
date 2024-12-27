package auth_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

func (delivery *AuthDeliveryImpl) Signup(c *gin.Context) {
	var req dto.ReqSignup

	// Parse json request
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Signup (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	res, jwtToken, httpCode, err := delivery.AuthUsecase.Signup(c.Request.Context(), req)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Signup (%s)", msg)
		c.String(httpCode, msg)
		return
	}

	c.SetCookie("jwt", jwtToken, (15 * 24 * 60 * 60), "/", "", false, true)
	c.JSON(httpCode, res)
}
