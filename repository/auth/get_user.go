package auth_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *AuthRepositoryImpl) GetUser(c context.Context, username string) (res entity.User, err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", c.Err().Error())
		log.Printf("Error in repository.AuthRepository.GetUser (%s)", msg)
		return entity.User{}, c.Err()
	}

	err = repository.DB.Collection("users").FindOne(c, bson.M{"userName": primitive.Regex{Pattern: fmt.Sprintf("^%s$", username), Options: "i"}}).Decode(&res)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", err.Error())
		log.Printf("Error in repository.AuthRepository.GetUser (%s)", msg)
		return entity.User{}, err
	}

	return res, nil
}
