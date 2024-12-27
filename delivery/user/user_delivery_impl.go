package user_delivery

import user_usecase "github.com/adhyttungga/go-chatapp-service/usecase/user"

type UserDeliveryImpl struct {
	UserUsecase user_usecase.UserUsecase
}

func NewUserDelivery(
	userUsecase user_usecase.UserUsecase,
) UserDelivery {
	return &UserDeliveryImpl{
		UserUsecase: userUsecase,
	}
}
