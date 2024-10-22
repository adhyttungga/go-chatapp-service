package message_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MessageRepositoryImpl) GetMessages(c context.Context, payload *entity.Conversation) (res []entity.Message, err error) {
	if c.Err() == context.DeadlineExceeded {
		msg := fmt.Sprintf("An error occured while retrieving an existing Message: %s", c.Err().Error())
		log.Printf("Error in message_repository.MessageRepository.GetMessages (%s)", msg)
		return []entity.Message{}, c.Err()
	}

	aggSearch := bson.M{"$match": bson.M{"participants": bson.M{"$all": payload.Participants}}}
	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "messages",
		"localField":   "messages",
		"foreignField": "_id",
		"as":           "listMessages",
	}}
	aggProject := bson.M{"$project": bson.M{"_id": 0, "participants": 0}}
	aggUnwind := bson.M{"$unwind": "$listMessages"}
	aggReplaceRoot := bson.M{"$replaceRoot": bson.M{"newRoot": "$listMessages"}}
	aggSort := bson.M{"$sort": bson.M{"createdAt": 1}}

	cursor, err := repository.DB.Collection("conversations").Aggregate(c, []bson.M{aggSearch, aggPopulate, aggProject, aggUnwind, aggReplaceRoot, aggSort})
	if err != nil {
		msg := fmt.Sprintf("An error occured while retrieving messages: %s", err.Error())
		log.Printf("Error in message_repository.MessageRepository.GetMessages (%s)", msg)
		return []entity.Message{}, err
	}

	var messages []entity.Message

	if err := cursor.All(c, &messages); err != nil {
		msg := fmt.Sprintf("An error occured while decoding the messages: %s", err.Error())
		log.Printf("Error in message_repository.MessageRepository.GetMessages (%s)", msg)
		return []entity.Message{}, err
	}

	return messages, nil
}
