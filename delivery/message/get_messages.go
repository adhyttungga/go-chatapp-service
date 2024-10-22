package message_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

func (delivery *MessageDeliveryImpl) GetMessages(c *gin.Context) {
	var req dto.ReqMessage
	senderID, _ := c.Get("user_id")
	req.SenderID = senderID.(string)
	req.ReceiverID = c.Param("id")

	res, err := delivery.MessageUsecase.GetMessages(c.Request.Context(), req)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while retrieving messages: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.GetMessages (%s)", msg)
		c.String(http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusOK, res)
}
