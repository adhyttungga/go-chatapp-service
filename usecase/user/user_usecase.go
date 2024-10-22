package user_usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	user_repository "github.com/adhyttungga/go-chatapp-service/repository/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	GetUser(c context.Context, loggedInUserID primitive.ObjectID) ([]dto.ResUser, error)
}

type UserUsecaseImpl struct {
	UserRepository user_repository.UserRepository
}

func NewUserUsecase(
	userRepository user_repository.UserRepository,
) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepository,
	}
}

func (usecase *UserUsecaseImpl) GetUser(c context.Context, loggedInUserID primitive.ObjectID) ([]dto.ResUser, error) {
	req := entity.User{
		ID: loggedInUserID,
	}

	res, err := usecase.UserRepository.GetUser(c, &req)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in user_usecase.UserUsecase.GetUser (%s)", msg) // Print error message to log
		return []dto.ResUser{}, err
	}

	var users []dto.ResUser
	for _, r := range res {
		users = append(users, dto.ResUser{
			ID:         r.ID.Hex(),
			FullName:   r.FullName,
			UserName:   r.UserName,
			ProfilePic: r.ProfilePic,
		})
	}

	fmt.Printf("\nUsers from usecase: \n%v", users)

	return users, nil
}
