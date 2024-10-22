package message_delivery

import (
	message_usecase "github.com/adhyttungga/go-chatapp-service/usecase/message"
	"github.com/go-playground/validator/v10"
)

type MessageDeliveryImpl struct {
	Validate       *validator.Validate
	MessageUsecase message_usecase.MessageUsecase
}

func NewMessageDelivery(
	messageUsecase message_usecase.MessageUsecase,
	validate *validator.Validate,
) MessageDelivery {
	return &MessageDeliveryImpl{
		MessageUsecase: messageUsecase,
		Validate:       validate,
	}
}
