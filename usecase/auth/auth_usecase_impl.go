package auth_usecase

import (
	user_repository "github.com/adhyttungga/go-chatapp-service/repository/user"
	"github.com/go-playground/validator/v10"
)

type AuthUsecaseImpl struct {
	UserRepository user_repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthUsecase(
	userRepository user_repository.UserRepository,
	validate *validator.Validate,
) AuthUsecase {
	return &AuthUsecaseImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
