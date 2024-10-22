package message_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *MessageRepositoryImpl) CreateMessage(c context.Context, payload *entity.Message) (resID primitive.ObjectID, err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while retrieving an existing conversation: %s", c.Err().Error())
		log.Printf("Error in repository.MessageRepository.CreateMessage (%s)", msg)
		return
	}

	response, err := repository.DB.Collection("messages").InsertOne(c, payload)
	if err != nil {
		msg := fmt.Sprintf("An error occured while creating message: %s", err.Error())
		log.Printf("Error in repository.MessageRepository.CreateMessage (%s)", msg)
		return resID, err
	}

	responseID, ok := response.InsertedID.(primitive.ObjectID)
	if !ok {
		err := errors.New("Cannot retrieve response ID")
		msg := fmt.Sprintf("An error occurred while retrieving response ID %s", err.Error())
		log.Printf("Error in repository.MessageRepository.CreateConversation (%s)", msg)
		return resID, err
	}

	return responseID, nil
}
