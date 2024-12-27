package message_repository

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
)

type MessageRepository interface {
	GetConversation(c context.Context, conversation *entity.Conversation) error
	CreateConversation(c context.Context, conversation *entity.Conversation) error
	UpdateConversation(c context.Context, conversation *entity.Conversation) error
	CreateMessage(c context.Context, message *entity.Message) error
	GetMessages(c context.Context, conversation *entity.Conversation) (messages []entity.Message, err error)
}
