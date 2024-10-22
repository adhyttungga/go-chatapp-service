package message_usecase

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
)

type MessageUsecase interface {
	SendMessage(c context.Context, req dto.ReqMessage) (dto.ResMessage, error)
	GetMessages(c context.Context, req dto.ReqMessage) ([]dto.ResMessage, error)
}
