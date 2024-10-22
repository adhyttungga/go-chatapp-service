package auth_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *AuthRepositoryImpl) CreateUser(c context.Context, payload *entity.User) (resID string, err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occurred while creating user: %s", c.Err().Error())
		log.Printf("Error in repository.AuthRepository.CreateUser (%s)", msg)
		return "", c.Err()
	}

	insertResult, err := repository.DB.Collection("users").InsertOne(c, payload)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in repository.AuthRepository.CreateUser (%s)", msg)
		return "", err
	}

	responseID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		err := errors.New("Cannot retrieve response ID")
		msg := fmt.Sprintf("An error occurred while retrieving response ID %s", err.Error())
		log.Printf("Error in repository.AuthRepository.CreateUser (%s)", msg)
		return "", err
	}

	return responseID.Hex(), err
}
