package message_usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (usecase *MessageUsecaseImpl) SendMessage(c context.Context, req dto.ReqMessage) (dto.ResMessage, error) {
	senderID, _ := primitive.ObjectIDFromHex(req.SenderID)
	receiverID, _ := primitive.ObjectIDFromHex(req.ReceiverID)
	var conversation entity.Conversation
	conversation.Participants = []primitive.ObjectID{senderID, receiverID}
	var newMessage entity.Message
	newMessage.Message = req.Message
	newMessage.SenderID = senderID
	newMessage.ReceiverID = receiverID

	res, err := usecase.MessageRepository.GetConversation(c, &conversation)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Print error message to log unless the error is "no document"
			msg := fmt.Sprintf("An error occured while retrieving an existing conversation: %s", err.Error())
			log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
			return dto.ResMessage{}, err
		}
	}

	conversation.ID = res.ID
	conversation.Messages = res.Messages
	conversation.CreatedAt = res.CreatedAt

	// Create new conversation if there is no an existing conversation
	if err == mongo.ErrNoDocuments {
		resID, err := usecase.MessageRepository.CreateConversation(c, &conversation)
		if err != nil {
			// Print error message to log unless the error is "no document"
			msg := fmt.Sprintf("An error occured while creating conversation: %s", err.Error())
			log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
			return dto.ResMessage{}, err
		}

		conversation.ID = resID
	}

	// Create new message
	resID, err := usecase.MessageRepository.CreateMessage(c, &newMessage)
	if err != nil {
		msg := fmt.Sprintf("An error occured while creating message: %s", err.Error())
		log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
		return dto.ResMessage{}, err
	}

	conversation.Messages = append(conversation.Messages, resID)
	log.Println(conversation)

	// Update conversation
	if err := usecase.MessageRepository.UpdateConversation(c, &conversation); err != nil {
		msg := fmt.Sprintf("An error occured while updating conversation: %s", err.Error())
		log.Printf("Error in message_usecase.MessageUsecase.SendMessage (%s)", msg)
		return dto.ResMessage{}, err
	}

	resMessage := dto.ResMessage{
		ID:         resID.Hex(),
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		Message:    req.Message,
	}

	return resMessage, nil
}
