package user_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *UserRepositoryImpl) Create(c context.Context, user *entity.User) (err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occurred while creating user: %s", c.Err().Error())
		log.Printf("Error in repository.UserRepository.Create (%s)", msg)
		return c.Err()
	}

	result, err := repository.DB.Collection("users").InsertOne(c, user)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in repository.UserRepository.Create (%s)", msg)
		return err
	}

	// Retrieving user ID
	userID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err := errors.New("Cannot retrieve user ID")
		msg := fmt.Sprintf("An error occurred while retrieving user ID %s", err.Error())
		log.Printf("Error in repository.UserRepository.Create (%s)", msg)
		return err
	}

	user.ID = userID
	return nil
}
