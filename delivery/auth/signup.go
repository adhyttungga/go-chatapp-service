package auth_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Validate the request
	if err := delivery.Validate.Struct(req); err != nil {
		// Return bad request and the message error
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Validate password and confirm password should match
	if req.Password != req.ConfirmPassword {
		// Return bad request and the message error if not match
		c.String(http.StatusBadRequest, "Password doesn't match")
		return
	}

	// Validate unique username
	err := delivery.AuthUsecase.GetUserByUsername(c.Request.Context(), req.UserName)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Print error message to log unless the error is no document found
			msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", err.Error())
			log.Printf("Error in delivery.AuthDelivery.Signup (%s)", msg)
			c.String(http.StatusInternalServerError, msg)
			return
		}
	} else {
		// Return bad request when username already exist
		c.String(http.StatusBadRequest, "Username already exists, please use another username")
		return
	}

	res, jwtToken, err := delivery.AuthUsecase.Signup(c.Request.Context(), req)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in delivery.AuthDelivery.Signup (%s)", msg)
		c.String(http.StatusInternalServerError, msg)
		return
	}

	c.SetCookie("jwt", jwtToken, (15 * 24 * 60 * 60), "/", "", false, true)
	c.JSON(http.StatusOK, res)
}
