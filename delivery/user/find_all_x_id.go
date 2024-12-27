package user_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (delivery *UserDeliveryImpl) FindAllExcludeId(c *gin.Context) {
	userID, _ := c.Get("user_id")
	loggedInUserID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while creating objectID from hex: %s", err.Error())
		log.Printf("Error in user_delivery.UserDelivery.FindAllExcludeId (%s)", msg) // Print error message to log
		c.String(http.StatusBadRequest, msg)
		return
	}

	users, err := delivery.UserUsecase.FindAllExcludeId(c, loggedInUserID)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while retrieving users: %s", err.Error())
		log.Printf("Error in user_delivery.UserDelivery.FindAllExcludeId (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, users)
}
