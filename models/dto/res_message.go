package dto

type ResMessage struct {
	ID         string `json:"_id"`
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Message    string `json:"message"`
}
