package user

import (
	"context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	InsertUser(
		ctx context.Context,
		user *model.User,
	) (*model.User, error)
}
