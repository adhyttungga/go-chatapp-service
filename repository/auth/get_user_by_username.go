package auth_repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *AuthRepositoryImpl) GetUserByUsername(c context.Context, username string) (err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", c.Err().Error())
		log.Printf("Error in repository.AuthRepository.GetUserByUsername (%s)", msg)
		return c.Err()
	}

	err = repository.DB.Collection("users").FindOne(c, bson.M{"userName": primitive.Regex{Pattern: fmt.Sprintf("^%s$", username), Options: "i"}}).Err()
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Print error message to log unless the error is no document found
			msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", err.Error())
			log.Printf("Error in repository.AuthRepository.GetUserByUsername (%s)", msg)
		}
	}

	return err
}
