package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SenderID   primitive.ObjectID `json:"senderId" bson:"senderId,omitempty"`
	ReceiverID primitive.ObjectID `json:"receiverId" bson:"receiverId,omitempty"`
	Message    string             `json:"message" bson:"message,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (u *Message) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my Message
	return bson.Marshal((*my)(u))
}
