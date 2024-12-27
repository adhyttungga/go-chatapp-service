package auth_usecase

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/dto"
)

type AuthUsecase interface {
	Signup(c context.Context, signup dto.ReqSignup) (resUser dto.ResUser, token string, httpCode int, err error)
	Login(c context.Context, login dto.ReqLogin) (httpCode int, resUser dto.ResUser, token string, err error)
}
