package message_usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (usecase *MessageUsecaseImpl) GetMessages(c context.Context, req dto.ReqMessage) ([]dto.ResMessage, error) {
	senderID, _ := primitive.ObjectIDFromHex(req.SenderID)
	receiverID, _ := primitive.ObjectIDFromHex(req.ReceiverID)
	var conversation entity.Conversation
	conversation.Participants = []primitive.ObjectID{senderID, receiverID}

	res, err := usecase.MessageRepository.GetMessages(c, &conversation)
	if err != nil {
		msg := fmt.Sprintf("An error occured while retrieving messages: %s", err.Error())
		log.Printf("Error in message_usecase.MessageUsecase.GetMessages (%s)", msg)
		return []dto.ResMessage{}, err
	}

	messages := []dto.ResMessage{}
	for _, r := range res {
		messages = append(messages, dto.ResMessage{
			ID:         r.ID.Hex(),
			SenderID:   r.SenderID.Hex(),
			ReceiverID: r.ReceiverID.Hex(),
			Message:    r.Message,
		})
	}

	return messages, nil
}
