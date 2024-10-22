package dto

type ReqMessage struct {
	Message    string `json:"message" validate:"required"`
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
}
