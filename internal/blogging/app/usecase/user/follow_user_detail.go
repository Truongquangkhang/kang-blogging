package user

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"kang-blogging/internal/blogging/domain/follow"
	"kang-blogging/internal/blogging/domain/user"
	"kang-blogging/internal/common/decorator"
	"kang-blogging/internal/common/errors"
	"kang-blogging/internal/common/model"
)

type FollowUserDetailParams struct {
	FollowedUserID string
	FollowerUserID string
}

type FollowUserDetailResult struct {
}

type FollowUserDetailHandler decorator.UsecaseHandler[FollowUserDetailParams, FollowUserDetailResult]

type followUserDetailHandler struct {
	userRepo   user.Repository
	followRepo follow.Repository
}

func NewFollowUserDetailHandler(
	userRepo user.Repository,
	followRepo follow.Repository,
	logger *logrus.Entry,
	metrics decorator.MetricsClient,
) FollowUserDetailHandler {
	if userRepo == nil {
		panic("blogRepo is nil")
	}
	if followRepo == nil {
		panic("followRepo is nil")
	}
	return decorator.ApplyUsecaseDecorators[FollowUserDetailParams, FollowUserDetailResult](
		followUserDetailHandler{
			userRepo:   userRepo,
			followRepo: followRepo,
		},
		logger,
		metrics,
	)
}

func (g followUserDetailHandler) Handle(ctx context.Context, param FollowUserDetailParams) (FollowUserDetailResult, error) {
	err := param.Validate()
	if err != nil {
		return FollowUserDetailResult{}, err
	}
	_, err = g.followRepo.CreateFollow(ctx, &model.Follow{
		FollowerID: param.FollowerUserID,
		FollowedID: param.FollowedUserID,
	})
	if err != nil {
		return FollowUserDetailResult{}, err
	}
	return FollowUserDetailResult{}, nil
}

func (p *FollowUserDetailParams) Validate() error {
	if p.FollowedUserID == "" || p.FollowerUserID == "" {
		return errors.NewBadRequestDefaultError()
	}
	return nil
}
