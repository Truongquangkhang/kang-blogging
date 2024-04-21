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

	GetUsers(
		ctx context.Context,
		params UserParams,
	) ([]model.User, int32, error)

	GetUserByID(
		ctx context.Context,
		id string,
	) (*model.User, error)

	UpdateUser(
		ctx context.Context,
		user *UserInfo,
	) (*model.User, error)
}
