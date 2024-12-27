package message_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

func (delivery *MessageDeliveryImpl) SendMessage(c *gin.Context) {
	var message dto.ReqMessage
	senderID, _ := c.Get("user_id")
	message.SenderID = senderID.(string)
	message.ReceiverID = c.Param("id")

	// Parse json request
	if err := c.ShouldBindJSON(&message); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.SendMessage (%s)", msg) // Print error message to log
		c.String(http.StatusInternalServerError, msg)
		return
	}

	resMessage, httpCode, err := delivery.MessageUsecase.SendMessage(c.Request.Context(), message)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while sending message: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.SendMessage (%s)", msg) // Print error message to log
		c.String(httpCode, msg)
		return
	}

	c.JSON(httpCode, resMessage)
}
