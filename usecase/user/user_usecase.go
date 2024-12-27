package user_usecase

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	FindAllExcludeId(c context.Context, loggedInUserID primitive.ObjectID) ([]dto.ResUser, error)
}
