package auth_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

// CreateTags		godoc
// @Summary			Login endpoint.
// @Accept			json
// @Produce			json
// @Tags				auth
// @Param				request body dto.ReqLogin true "Request body"
// @Success			200 {object} dto.ResUser "Ok"
// @Failure			400 {object} dto.ResError "Bad request error"
// @Failure			500 {object} dto.ResError "Internal server error"
// @Response    default {object} dto.ResError "Other error"
// @Router			/auth/login [post]
func (delivery *AuthDeliveryImpl) Login(c *gin.Context) {
	var login dto.ReqLogin
	var resError dto.ResError

	// Parse json request
	if err := c.ShouldBindJSON(&login); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Login (%s)", msg) // Print error message to log
		resError.Error = msg
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	httpCode, user, jwtToken, err := delivery.AuthUsecase.Login(c.Request.Context(), login)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Login (%s)", msg)
		resError.Error = msg
		c.JSON(httpCode, resError)
		return
	}

	c.SetCookie("jwt", jwtToken, (15 * 24 * 60 * 60), "/", "", false, true)
	c.JSON(httpCode, user)
}
