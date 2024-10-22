package message_delivery

import (
	"github.com/gin-gonic/gin"
)

type MessageDelivery interface {
	SendMessage(c *gin.Context)
	GetMessages(c *gin.Context)
}
