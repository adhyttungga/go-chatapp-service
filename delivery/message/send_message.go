package message_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

// CreateTags		godoc
// @Summary			Send message endpoint.
// @Description Send messages from logged in user to other user.
// @Accept			json
// @Produce			json
// @Tags				message
// @Param				id path string true "Id"
// @Param				request body dto.ReqMessage true "Request body"
// @Success			201 {object} dto.ResMessage "Created"
// @Failure			400 {object} dto.ResError "Bad request error"
// @Failure			401 {object} dto.ResError "Unauthorized error"
// @Failure			500 {object} dto.ResError "Internal server error"
// @Response		default {object} dto.ResError "Other error"
// @Router			/message/send/{id} [post]
func (delivery *MessageDeliveryImpl) SendMessage(c *gin.Context) {
	var (
		message  dto.ReqMessage
		resError dto.ResError
	)

	// Parse json request
	if err := c.ShouldBindJSON(&message); err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while parsing json request: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.SendMessage (%s)", msg) // Print error message to log
		resError.Error = msg
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	senderID, _ := c.Get("user_id")
	message.SenderID = senderID.(string)
	message.ReceiverID = c.Param("id")

	resMessage, httpCode, err := delivery.MessageUsecase.SendMessage(c.Request.Context(), message)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while sending message: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.SendMessage (%s)", msg) // Print error message to log
		resError.Error = msg
		c.JSON(httpCode, resError)
		return
	}

	c.JSON(httpCode, resMessage)
}
