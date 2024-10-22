package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FullName   string             `json:"fullName" bson:"fullName,omitempty"`
	UserName   string             `json:"userName" bson:"userName,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	Gender     string             `json:"gender" bson:"gender,omitempty"`
	ProfilePic string             `json:"profilePic" bson:"profilePic,omitempty"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (u *User) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my User
	return bson.Marshal((*my)(u))
}
