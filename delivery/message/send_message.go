package message_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

func (delivery *MessageDeliveryImpl) SendMessage(c *gin.Context) {
	var req dto.ReqMessage
	senderID, _ := c.Get("user_id")
	req.SenderID = senderID.(string)
	req.ReceiverID = c.Param("id")

	// Parse json request
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.SendMessage (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	// Validate the request
	if err := delivery.Validate.Struct(req); err != nil {
		// Return bad request and the message error
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resMessage, err := delivery.MessageUsecase.SendMessage(c.Request.Context(), req)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while sending message: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.SendMessage (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusCreated, resMessage)
}
