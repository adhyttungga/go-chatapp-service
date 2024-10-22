package message_repository

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository interface {
	GetConversation(c context.Context, payload *entity.Conversation) (res entity.Conversation, err error)
	CreateConversation(c context.Context, payload *entity.Conversation) (resID primitive.ObjectID, err error)
	UpdateConversation(c context.Context, payload *entity.Conversation) error
	CreateMessage(c context.Context, payload *entity.Message) (resID primitive.ObjectID, err error)
	GetMessages(c context.Context, payload *entity.Conversation) (res []entity.Message, err error)
}
