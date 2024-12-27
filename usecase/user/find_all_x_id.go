package user_usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (usecase *UserUsecaseImpl) FindAllExcludeId(c context.Context, loggedInUserID primitive.ObjectID) ([]dto.ResUser, error) {
	userEntity := entity.User{
		ID: loggedInUserID,
	}

	// Retrieving all users exclude logged in user
	usersEntity, err := usecase.UserRepository.FindAllExcludeId(c, &userEntity)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in user_usecase.UserUsecase.FindAllExcludeId (%s)", msg) // Print error message to log
		return []dto.ResUser{}, err
	}

	var users []dto.ResUser
	for _, r := range usersEntity {
		users = append(users, dto.ResUser{
			ID:         r.ID.Hex(),
			FullName:   r.FullName,
			UserName:   r.UserName,
			ProfilePic: r.ProfilePic,
		})
	}

	return users, nil
}
