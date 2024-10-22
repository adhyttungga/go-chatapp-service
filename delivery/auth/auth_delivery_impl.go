package auth_delivery

import (
	auth_usecase "github.com/adhyttungga/go-chatapp-service/usecase/auth"
	"github.com/go-playground/validator/v10"
)

type AuthDeliveryImpl struct {
	Validate    *validator.Validate
	AuthUsecase auth_usecase.AuthUsecase
}

func NewAuthDelivery(authUsecase auth_usecase.AuthUsecase, validate *validator.Validate) AuthDelivery {
	return &AuthDeliveryImpl{
		AuthUsecase: authUsecase,
		Validate:    validate,
	}
}
