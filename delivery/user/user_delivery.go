package user_delivery

import (
	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	FindAllExcludeId(c *gin.Context)
}
