package message_delivery

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/gin-gonic/gin"
)

// CreateTags		godoc
// @Summary			Get message endpoint.
// @Description List messages from logged in user with other user.
// @Accept			json
// @Produce			json
// @Tags				message
// @Param				id path string true "Id"
// @Success			200 {array} dto.ResMessage "Ok"
// @Failure			401 {object} dto.ResError "Unauthorized error"
// @Failure			500 {object} dto.ResError "Internal server error"
// @Response		default {object} dto.ResError "Other error"
// @Router			/message/{id} [get]
func (delivery *MessageDeliveryImpl) GetMessages(c *gin.Context) {
	var (
		resError dto.ResError
		req      dto.ReqMessage
	)
	senderID, _ := c.Get("user_id")
	req.SenderID = senderID.(string)
	req.ReceiverID = c.Param("id")

	res, err := delivery.MessageUsecase.GetMessages(c.Request.Context(), req)
	if err != nil {
		// Return internal server error and the message error
		msg := fmt.Sprintf("An error occurred while retrieving messages: %s", err.Error())
		log.Printf("Error in delivery.MessageDelivery.GetMessages (%s)", msg)
		resError.Error = msg
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	c.JSON(http.StatusOK, res)
}
