package message_repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *MessageRepositoryImpl) CreateConversation(c context.Context, conversation *entity.Conversation) error {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while retrieving an existing conversation: %s", c.Err().Error())
		log.Printf("Error in repository.MessageRepository.CreateConversation (%s)", msg)
		return c.Err()
	}

	result, err := repository.DB.Collection("conversations").InsertOne(c, conversation)
	if err != nil {
		msg := fmt.Sprintf("An error occured while creating conversation: %s", err.Error())
		log.Printf("Error in repository.MessageRepository.CreateConversation (%s)", msg)
		return err
	}

	conversationID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err := errors.New("Cannot retrieve conversation ID")
		msg := fmt.Sprintf("An error occurred while retrieving conversation ID %s", err.Error())
		log.Printf("Error in repository.MessageRepository.CreateConversation (%s)", msg)
		return err
	}

	conversation.ID = conversationID
	return nil
}
