package user_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTags		godoc
// @Summary			Find all user exclude logged in user endpoint.
// @Description	Displays all users except the currently logged in user.
// @Accept			json
// @Produce			json
// @Tags				user
// @Success			200 {array} dto.ResUser "Ok"
// @Failure			401 {object} dto.ResError "Unauthorized error"
// @Failure			500 {object} dto.ResError "Internal server error"
// @Response    default {object} dto.ResError "Other error"
// @Router			/user/ [get]
func (delivery *UserDeliveryImpl) FindAllExcludeId(c *gin.Context) {
	var resError dto.ResError
	userID, _ := c.Get("user_id")
	loggedInUserID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while creating objectID from hex: %s", err.Error())
		log.Printf("Error in user_delivery.UserDelivery.FindAllExcludeId (%s)", msg)
		resError.Error = msg
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	users, err := delivery.UserUsecase.FindAllExcludeId(c, loggedInUserID)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while retrieving users: %s", err.Error())
		log.Printf("Error in user_delivery.UserDelivery.FindAllExcludeId (%s)", msg)
		resError.Error = msg
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	c.JSON(http.StatusOK, users)
}
