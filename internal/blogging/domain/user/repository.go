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
		user *model.User,
	) (*model.User, error)

	GetInfoFromMultiTable(
		ctx context.Context,
	) (*SystemInfo, error)

	ChangeStatus(
		ctx context.Context,
		userId string,
		currentStatus bool,
	) error

	GetRelateInfoOfUser(
		ctx context.Context,
		userId string,
		ignoreBlogIsDraft bool,
		currentUserId *string,
	) (*RelateUserInfo, error)
}
