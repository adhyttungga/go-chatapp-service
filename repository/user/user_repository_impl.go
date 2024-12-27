package user_repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepositoryImpl struct {
	DB *mongo.Database
}

func NewUserRepository(DB *mongo.Database) UserRepository {
	return &UserRepositoryImpl{DB}
}
