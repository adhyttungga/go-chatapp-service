package user_delivery

import (
	"fmt"
	"log"
	"net/http"

	user_usecase "github.com/adhyttungga/go-chatapp-service/usecase/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDelivery interface {
	GetUser(c *gin.Context)
}

type UserDeliveryImpl struct {
	UserUsecase user_usecase.UserUsecase
}

func NewUserDelivery(
	userUsecase user_usecase.UserUsecase,
) UserDelivery {
	return &UserDeliveryImpl{
		UserUsecase: userUsecase,
	}
}

func (delivery *UserDeliveryImpl) GetUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	loggedInUserID, _ := primitive.ObjectIDFromHex(userID.(string))

	res, err := delivery.UserUsecase.GetUser(c, loggedInUserID)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in user_delivery.UserDelivery.GetUser (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, res)
}
