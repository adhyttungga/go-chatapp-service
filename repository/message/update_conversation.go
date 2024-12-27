package message_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MessageRepositoryImpl) UpdateConversation(c context.Context, conversation *entity.Conversation) error {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while updating conversation: %s", c.Err().Error())
		log.Printf("Error in repository.MessageRepository.UpdateConversation (%s)", msg)
		return c.Err()
	}

	_, err := repository.DB.Collection("conversations").UpdateOne(c, bson.M{"_id": conversation.ID}, bson.M{"$set": conversation})
	if err != nil {
		msg := fmt.Sprintf("An error occured while updating conversation: %s", err.Error())
		log.Printf("Error in repository.MessageRepository.UpdateConversation (%s)", msg)
		return err
	}

	return nil
}
