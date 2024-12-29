package message_usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (usecase *MessageUsecaseImpl) SendMessage(c context.Context, req dto.ReqMessage) (dto.ResMessage, int, error) {
	emptyMessage := dto.ResMessage{}

	// Validate the request
	if err := usecase.Validate.Struct(req); err != nil {
		// Return bad request and the message error
		return emptyMessage, http.StatusBadRequest, err
	}

	senderID, _ := primitive.ObjectIDFromHex(req.SenderID)
	receiverID, _ := primitive.ObjectIDFromHex(req.ReceiverID)
	var conversation entity.Conversation
	conversation.Participants = []primitive.ObjectID{senderID, receiverID}
	var newMessage entity.Message
	newMessage.Message = req.Message
	newMessage.SenderID = senderID
	newMessage.ReceiverID = receiverID

	err := usecase.MessageRepository.GetConversation(c, &conversation)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Print error message to log unless the error is "no document"
			msg := fmt.Sprintf("An error occured while retrieving an existing conversation: %s", err.Error())
			log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
			return emptyMessage, http.StatusInternalServerError, err
		} else { // Create new conversation if there is no an existing conversation
			if err := usecase.MessageRepository.CreateConversation(c, &conversation); err != nil {
				msg := fmt.Sprintf("An error occured while creating conversation: %s", err.Error())
				log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
				return emptyMessage, http.StatusInternalServerError, err
			}
		}
	}

	// Create new message
	if err := usecase.MessageRepository.CreateMessage(c, &newMessage); err != nil {
		msg := fmt.Sprintf("An error occured while creating message: %s", err.Error())
		log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
		return emptyMessage, http.StatusInternalServerError, err
	}

	// Update conversation
	conversation.Messages = append(conversation.Messages, newMessage.ID)
	if err := usecase.MessageRepository.UpdateConversation(c, &conversation); err != nil {
		msg := fmt.Sprintf("An error occured while updating conversation: %s", err.Error())
		log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
		return emptyMessage, http.StatusInternalServerError, err
	}

	resMessage := dto.ResMessage{
		ID:         newMessage.ID.Hex(),
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		Message:    req.Message,
	}

	return resMessage, http.StatusCreated, nil
}
