package auth_delivery

import (
	auth_usecase "github.com/adhyttungga/go-chatapp-service/usecase/auth"
)

type AuthDeliveryImpl struct {
	AuthUsecase auth_usecase.AuthUsecase
}

func NewAuthDelivery(authUsecase auth_usecase.AuthUsecase) AuthDelivery {
	return &AuthDeliveryImpl{
		AuthUsecase: authUsecase,
	}
}
