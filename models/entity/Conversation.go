package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	ID           primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Participants []primitive.ObjectID `json:"participants" bson:"participants,omitempty"`
	Messages     []primitive.ObjectID `json:"messages" bson:"messages,omitempty"`
	CreatedAt    time.Time            `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time            `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (u *Conversation) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my Conversation
	return bson.Marshal((*my)(u))
}
