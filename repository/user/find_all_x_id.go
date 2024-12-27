package user_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *UserRepositoryImpl) FindAllExcludeId(c context.Context, userEntity *entity.User) ([]entity.User, error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while retrieving users: %s", c.Err().Error())
		log.Printf("Error in user_repository.UserRepository.FindAllExcludeId (%s)", msg)
		return []entity.User{}, c.Err()
	}

	filter := bson.M{"_id": bson.M{"$ne": userEntity.ID}}
	cursor, err := repository.DB.Collection("users").Find(c, filter)
	if err != nil {
		// Print error message to log unless the error is no document found
		msg := fmt.Sprintf("An error occurred while retrieving users: %s", err.Error())
		log.Printf("Error in user_repository.UserRepository.FindAllExcludeId (%s)", msg)
		return []entity.User{}, err
	}

	var users []entity.User

	if err := cursor.All(c, &users); err != nil {
		msg := fmt.Sprintf("An error occured while decoding the users: %s", err.Error())
		log.Printf("Error in user_repository.UserRepository.FindAllExcludeId (%s)", msg)
		return []entity.User{}, err
	}

	return users, nil
}
