package auth_usecase

import (
	auth_repository "github.com/adhyttungga/go-chatapp-service/repository/auth"
)

type AuthUsecaseImpl struct {
	AuthRepository auth_repository.AuthRepository
}

func NewAuthUsecase(authRepository auth_repository.AuthRepository) AuthUsecase {
	return &AuthUsecaseImpl{
		AuthRepository: authRepository,
	}
}
