package user_repository

import (
	"context"

	"github.com/adhyttungga/go-chatapp-service/models/entity"
)

type UserRepository interface {
	Create(c context.Context, user *entity.User) (err error)
	FindByUsername(c context.Context, user *entity.User) (err error)
	FindAllExcludeId(c context.Context, userEntity *entity.User) ([]entity.User, error)
}
