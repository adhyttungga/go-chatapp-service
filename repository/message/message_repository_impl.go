package message_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepositoryImpl struct {
	DB *mongo.Database
}

func NewMessageRepository(DB *mongo.Database) MessageRepository {
	return &MessageRepositoryImpl{DB}
}
