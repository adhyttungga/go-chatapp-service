package user_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *UserRepositoryImpl) FindByUsername(c context.Context, user *entity.User) (err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occurred while retrieving an existing user by username: %s", c.Err().Error())
		log.Printf("Error in repository.UserRepository.FindByUsername (%s)", msg)
		return c.Err()
	}

	if err := repository.DB.Collection("users").FindOne(c, bson.M{"userName": primitive.Regex{Pattern: fmt.Sprintf("^%s$", user.UserName), Options: "i"}}).Decode(&user); err != nil {
		msg := fmt.Sprintf("An error occurred while retrieving an existing user by username: %s", err.Error())
		log.Printf("Error in repository.UserRepository.FindByUsername (%s)", msg)
		return err
	}

	return nil
}
