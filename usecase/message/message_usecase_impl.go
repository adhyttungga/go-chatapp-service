package message_usecase

import (
	message_repository "github.com/adhyttungga/go-chatapp-service/repository/message"
)

type MessageUsecaseImpl struct {
	MessageRepository message_repository.MessageRepository
}

func NewMessageUsecase(
	messageRepository message_repository.MessageRepository,
) MessageUsecase {
	return &MessageUsecaseImpl{
		MessageRepository: messageRepository,
	}
}
