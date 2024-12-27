package message_usecase

import (
	message_repository "github.com/adhyttungga/go-chatapp-service/repository/message"
	"github.com/go-playground/validator/v10"
)

type MessageUsecaseImpl struct {
	MessageRepository message_repository.MessageRepository
	Validate          *validator.Validate
}

func NewMessageUsecase(
	messageRepository message_repository.MessageRepository,
	validate *validator.Validate,
) MessageUsecase {
	return &MessageUsecaseImpl{
		MessageRepository: messageRepository,
		Validate:          validate,
	}
}
