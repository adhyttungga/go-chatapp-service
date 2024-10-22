package auth_repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryImpl struct {
	DB *mongo.Database
}

func NewAuthRepository(
	DB *mongo.Database) AuthRepository {
	return &AuthRepositoryImpl{DB}
}
