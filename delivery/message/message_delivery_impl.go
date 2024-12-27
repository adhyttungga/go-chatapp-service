package message_delivery

import (
	message_usecase "github.com/adhyttungga/go-chatapp-service/usecase/message"
)

type MessageDeliveryImpl struct {
	MessageUsecase message_usecase.MessageUsecase
}

func NewMessageDelivery(
	messageUsecase message_usecase.MessageUsecase,
) MessageDelivery {
	return &MessageDeliveryImpl{
		MessageUsecase: messageUsecase,
	}
}
