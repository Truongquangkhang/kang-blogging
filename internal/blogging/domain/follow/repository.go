package follow

import (
	"golang.org/x/net/context"
	"kang-blogging/internal/common/model"
)

type Repository interface {
	CreateFollow(
		ctx context.Context,
		follow *model.Follow,
	) (*model.Follow, error)

	DeleteFollow(
		ctx context.Context,
		follow *model.Follow,
	) (*model.Follow, error)

	GetFollowedIdsByFollowerId(
		ctx context.Context,
		followerId string,
	) ([]string, error)
}
