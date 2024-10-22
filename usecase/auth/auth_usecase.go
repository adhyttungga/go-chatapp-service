package auth_usecase

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
)

type AuthUsecase interface {
	Signup(c context.Context, req dto.ReqSignup) (res dto.ResUser, token string, err error)
	GetUserByUsername(c context.Context, username string) (err error)
	Login(c context.Context, req dto.ReqLogin) (httpCode int, res dto.ResUser, token string, err error)
}
