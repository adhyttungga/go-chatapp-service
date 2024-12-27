package user_usecase

import (
	user_repository "github.com/adhyttungga/go-chatapp-service/repository/user"
)

type UserUsecaseImpl struct {
	UserRepository user_repository.UserRepository
}

func NewUserUsecase(
	userRepository user_repository.UserRepository,
) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepository,
	}
}
