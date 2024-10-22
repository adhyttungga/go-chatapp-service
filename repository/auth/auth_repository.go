package auth_repository

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
)

type AuthRepository interface {
	CreateUser(context.Context, *entity.User) (resId string, err error)
	GetUserByUsername(c context.Context, username string) (err error)
	GetUser(c context.Context, username string) (res entity.User, err error)
}
