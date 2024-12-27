package message_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MessageRepositoryImpl) GetConversation(c context.Context, conversation *entity.Conversation) error {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while retrieving an existing conversation: %s", c.Err().Error())
		log.Printf("Error in message_repository.MessageRepository.GetConversation (%s)", msg)
		return c.Err()
	}

	if err := repository.DB.Collection("conversations").FindOne(c, bson.M{"participants": bson.M{"$all": conversation.Participants}}).Decode(&conversation); err != nil {
		// Print error message to log unless the error is no document found
		msg := fmt.Sprintf("An error occurred while retrieving an existing conversation: %s", err.Error())
		log.Printf("Error in message_repository.MessageRepository.GetConversation (%s)", msg)
		return err
	}

	return nil
}
