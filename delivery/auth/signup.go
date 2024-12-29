package auth_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

// CreateTags		godoc
// @Summary			SignUp endpoint.
// @Accept			json
// @Produce			json
// @Tags				auth
// @Param				request body dto.ReqSignup true "Request body"
// @Success			201 {object} dto.ResUser "Created"
// @Failure			400 {object} dto.ResError "Bad request error"
// @Failure			500 {object} dto.ResError "Internal server error"
// @Response		default {object} dto.ResError "Other error"
// @Router			/auth/signup [post]
func (delivery *AuthDeliveryImpl) Signup(c *gin.Context) {
	var req dto.ReqSignup
	var resError dto.ResError

	// Parse json request
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Signup (%s)", msg) // Print error message to log
		resError.Error = msg
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	res, jwtToken, httpCode, err := delivery.AuthUsecase.Signup(c.Request.Context(), req)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Signup (%s)", msg)
		resError.Error = msg
		c.JSON(httpCode, resError)
		return
	}

	c.SetCookie("jwt", jwtToken, (15 * 24 * 60 * 60), "/", "", false, true)
	c.JSON(httpCode, res)
}
