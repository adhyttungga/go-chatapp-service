package message_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *MessageRepositoryImpl) CreateMessage(c context.Context, message *entity.Message) error {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while creating message: %s", c.Err().Error())
		log.Printf("Error in repository.MessageRepository.CreateMessage (%s)", msg)
		return c.Err()
	}

	result, err := repository.DB.Collection("messages").InsertOne(c, message)
	if err != nil {
		msg := fmt.Sprintf("An error occured while creating message: %s", err.Error())
		log.Printf("Error in repository.MessageRepository.CreateMessage (%s)", msg)
		return err
	}

	messageID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err := errors.New("Cannot retrieve message ID")
		msg := fmt.Sprintf("An error occurred while retrieving message ID %s", err.Error())
		log.Printf("Error in repository.MessageRepository.CreateMessage (%s)", msg)
		return err
	}

	message.ID = messageID
	return nil
}
